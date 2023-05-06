package main

import (
	"GoAsk/cache"
	"GoAsk/config"
	"GoAsk/dao"
	"GoAsk/routes"
	"strings"
)

func main() {
	// 配置初始化
	config.Init("./config/config.ini")
	srcDSN := strings.Join([]string{config.DbUser, ":", config.DbPassWord, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	repDSN := strings.Join([]string{config.DbUser, ":", config.DbPassWord, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	db := dao.ConnectWithReplica(srcDSN, repDSN)
	// mysql 数据库初始化
	dao.Init(db)
	// redis 缓存初始化
	cache.Init()
	r := routes.NewRouter()
	if err := r.Run(config.HttpPort); err != nil {
		panic(err)
	}
}
