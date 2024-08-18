package main

import (
	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/db"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
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
	// 初始化数据库连接，设置连接池数量
	db.InitDatabaseConnectionPool()
	// 检测数据库结构是否正确，自动处理数据库中表的缺失/缺列等问题
	db.MigrateTables()

	// 启动Bot前发送一个初始化成功事件
	err = core.SendInitializeEvent()
	if err != nil {
		return
	}

	err = core.CreateUser(2913844577)
	if err != nil {
		return
	}
	// 启动Bot
	zerobot.LaunchQQBot()
}
