package plugins

import (
	"fmt"

	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type FetchPlugin struct{}

func (p *FetchPlugin) PluginInfos() (string, string, string, string) {
	var pluginId = "iceinu-fetch-plugin"     // 插件id，不能重复
	var pluginVersion = "1.0.0"              // 插件版本
	var pluginIntroduction = "用于展示系统运行信息的插件" // 插件简介
	var pluginDeveloper = "Kyoku"            // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *FetchPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnCommand("fetch").Handle(func(ctx *zero.Ctx) {
		fetchInfo := core.GetFetch()
		fetchText := fmt.Sprintf("Iceinu 运行信息\n>>>\n处理器：%s\n核心数：%s\n运行频率：%s\n总内存：%s\n系统类型：%s\n系统名称：%s\n内核版本：%s", fetchInfo.CPU, fetchInfo.Cores, fetchInfo.Frequency, fetchInfo.Memory, fetchInfo.SystemType, fetchInfo.Platform, fetchInfo.KernelVersion)
		ctx.SendChain(message.Text(fetchText))
	})
}

func (p *FetchPlugin) PluginHelps() []models.CommandHelp {
	return []models.CommandHelp{
		{
			IsShown:     true,
			CommandName: "fetch",
			Usage:       fmt.Sprintf("%sfetch", cfg.Bot.CommandPrefix),
			Description: "显示系统运行信息",
			Flags:       map[string]string{},
		},
	}
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&FetchPlugin{})
}
