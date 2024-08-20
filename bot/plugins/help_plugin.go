package plugins

import (
	"fmt"

	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type HelpPlugin struct{}

var cfg, _ = config.GetConfig()

func (p *HelpPlugin) PluginInfos() (string, string, string, string) {
	var pluginId = "iceinu-help-plugin"    // 插件id，不能重复
	var pluginVersion = "0.0.1"            // 插件版本
	var pluginIntroduction = "用于展示帮助信息的插件" // 插件简介
	var pluginDeveloper = "Kyoku"          // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *HelpPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnCommandGroup([]string{"help", "帮助"}).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text(core.GetHelpList()))
	})
}

func (p *HelpPlugin) PluginHelps() []models.CommandHelp {
	return []models.CommandHelp{
		{
			IsShown:     true,
			CommandName: "help",
			Usage:       fmt.Sprintf("%shelp {命令名称 | 页数}", cfg.Bot.CommandPrefix),
			Description: "显示Bot的Help信息",
			Flags:       map[string]string{},
		},
	}
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&HelpPlugin{})
}
