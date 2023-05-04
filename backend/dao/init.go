package dao

import (
	"GoAsk/config"
	"GoAsk/model"
	"context"
	"gorm.io/plugin/dbresolver"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	_db *gorm.DB
)

type DbClient struct {
	*gorm.DB
}

func NewDbClient(ctx context.Context) *DbClient {
	// https://gorm.io/docs/method_chaining.html WithContext 中调用了.Session(), 即自动从连接池中获取新的连接, 是并发安全的
	return &DbClient{
		_db.WithContext(ctx),
	}
}

func Init() {
	srcDSN := strings.Join([]string{config.DbUser, ":", config.DbPassWord, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	repDSN := strings.Join([]string{config.DbUser, ":", config.DbPassWord, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	ConnectDB(srcDSN, repDSN)
}

func ConnectDB(sourceDSN, replicaDSN string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       sourceDSN, // DSN data source name
		DefaultStringSize:         256,       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,     // 根据版本自动配置
	}), &gorm.Config{
		Logger:         ormLogger,
		NamingStrategy: schema.NamingStrategy{
			//SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	if err := db.Use(
		dbresolver.Register(dbresolver.Config{ // 主从库设置
			Replicas: []gorm.Dialector{mysql.Open(replicaDSN)},
		}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	); err != nil {
		panic(err)
	}
	_db = db
	autoMigration()

}

func autoMigration() {
	if err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{}, &model.Like{}); err != nil {
		log.Fatalln("自动迁移失败", err)
	}
	log.Println("自动迁移成功")
}
