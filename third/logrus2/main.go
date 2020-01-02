package main

import (
	"github.com/sirupsen/logrus"
	"os"
)
func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.WarnLevel)
}

var log = logrus.New()

func main() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")

	log.Info("admin....")
	log.Info("java")
	log.Errorln("show")
	log.Errorln("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}
