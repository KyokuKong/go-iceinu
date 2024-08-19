package plugins

import (
	"fmt"

	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type ExamplePlugin struct{}

func (p *ExamplePlugin) PluginInfos() (string, string, string, string) {
	var pluginId = "iceinu-example-plugin"              // 插件id，不能重复
	var pluginVersion = "0.0.1"                         // 插件版本
	var pluginIntroduction = "这是iceinu bot的示例插件，没有任何功能" // 插件简介
	var pluginDeveloper = "Kyoku"                       // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *ExamplePlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnFullMatch("go-iceinu").Handle(func(ctx *zero.Ctx) {
		fl := ctx.GetFriendList()
		fmt.Println(fl.String())
		ctx.SendChain(message.Text("Powered by Go-Iceinu, star the project on Github!\nhttps://github.com/kyokukong/go-iceinu"))
	})
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&ExamplePlugin{})
}
