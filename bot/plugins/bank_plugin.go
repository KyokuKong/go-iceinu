package plugins

import (
	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type BankPlugin struct{}

func (p *BankPlugin) PluginInfos() (string, string, string, string) {
	var pluginId = "iceinu-bank-plugin"          // 插件id，不能重复
	var pluginVersion = "1.0.0"                  // 插件版本
	var pluginIntroduction = "用于实现Bot的经济管理系统的插件" // 插件简介
	var pluginDeveloper = "Kyoku"                // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *BankPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnCommand("go-iceinu").Handle(func(ctx *zero.Ctx) {
		// core.PrintHelpList()
		ctx.SendChain(message.Text("Powered by Go-Iceinu, star the project on Github!\nhttps://github.com/KyokuKong/go-iceinu"))
	})
}

func (p *BankPlugin) PluginHelps() []models.CommandHelp {
	return []models.CommandHelp{
		{
			IsShown:     false,
			CommandName: "go-iceinu",
			Usage:       "go-iceinu",
			Description: "触发冰犬 bot 的调试信息",
			Flags: map[string]string{
				"-v": "显示版本信息",
			},
		},
	}
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&BankPlugin{})
}
