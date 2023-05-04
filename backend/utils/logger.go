package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

func init() {
	if Logger != nil {
		src, _ := setOutputFile()
		//设置输出
		Logger.Out = src
		return
	}

	//实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	//设置输出
	logger.Out = src
	//设置日志级别, 一共有7级, 详见 https://github.com/sirupsen/logrus
	// Only log the warning DebugLevel or above.
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger = logger
}

func setOutputFile() (*os.File, error) { // 设置日志文件
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}

	_, err := os.Stat(logFilePath) // os.Stat 获取文件或者文件夹的信息
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil { // 创建log文件夹
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log" // time.now.Format 将time转为string格式
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil { // 不存在该.log文件则创建
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
