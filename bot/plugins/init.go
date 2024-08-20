package plugins

import (
	"github.com/KyokuKong/go-iceinu/bot/utils"
	log "github.com/sirupsen/logrus"
)

// 这不是一个插件，仅仅是一个用于触发plugins包初始化的工具函数
// 插件管理器会在注册插件之前运行这个函数来触发所有内置插件的init()函数

func InitPlugins() {
	log.Debug(utils.GenerateGradientString("\"Talk is cheap, Show me the code.\"", "#a1c4fd", "#c2e9fb"))
}
