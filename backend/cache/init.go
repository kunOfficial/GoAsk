package cache

import (
	"GoAsk/config"
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
)

type CacheClient struct {
	*redis.Client
}

func Init() {
	connectRedis()
	s := gocron.NewScheduler(time.UTC)
	// 2s 定时将 redis 中的浏览累积量同步到 mysql 中
	_, err := s.Every(3).Seconds().Do(func() { syncViews() })
	if err != nil {
		panic(err)
	}
	_, err = s.Every(2).Seconds().Do(func() { syncLikes() })
	//if err != nil {
	//	panic(err)
	//}
	s.StartAsync()
}

func NewCacheClient(ctx context.Context) *CacheClient {
	return &CacheClient{
		RedisClient.WithContext(ctx),
	}
}

//connectRedis 在中间件中初始化redis链接
func connectRedis() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
