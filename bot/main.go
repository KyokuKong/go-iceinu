package main

import (
	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/db"
	"github.com/KyokuKong/go-iceinu/bot/plugins"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 程序主函数，在这里运行整个初始化以及启动流程
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

	// 注册插件管理器
	plugins.InitPlugins()
	core.RegisterManager()

	// 启动Bot前发送一个初始化成功事件
	err = core.SendInitializeEvent()
	if err != nil {
		return
	}

	// 启动Bot
	zerobot.LaunchQQBot()
}
