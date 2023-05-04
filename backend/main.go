package main

import (
	"GoAsk/cache"
	"GoAsk/config"
	"GoAsk/dao"
	"GoAsk/routes"
)

func main() {
	// 配置初始化
	config.Init()
	// mysql 数据库初始化
	dao.Init()
	// redis 缓存初始化
	cache.Init()
	r := routes.NewRouter()
	if err := r.Run(config.HttpPort); err != nil {
		panic(err)
	}
}
