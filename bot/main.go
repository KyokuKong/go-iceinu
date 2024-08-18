package main

import (
	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/core"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {
	// 程序主函数，在这里运行整个初始化以及启动流程
	// 初始化日志管理模块
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.InfoLevel)
	// 初始化配置模块
	cfg, err := config.GetConfig()
	if err != nil {
		return
	}
	var logLevel = cfg.Bot.LogLevel
	// 根据配置重新设置日志等级
	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	}

	// 启动Bot
	core.LaunchQQBot()
}
