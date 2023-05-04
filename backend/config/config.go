package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
	AvatarPath  string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalln("配置文件读取失败:", err)
	}
	loadServer(file)
	loadMysql(file)
	loadRedis(file)
	loadPath(file)
}

func loadPath(file *ini.File) {
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}

// 加载服务器配置信息
func loadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// 加载mysql配置信息
func loadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

// 加载redis配置信息
func loadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
