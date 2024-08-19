package plugins

import (
	"fmt"

	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type ExamplePlugin struct{}

func (p *ExamplePlugin) PluginInfos() (string, string, string) {
	return "iceinu-example-plugin", "0.0.1", "这是iceinu bot的示例插件，没有任何功能"
}

func (p *ExamplePlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnFullMatch("iceinu").Handle(func(ctx *zero.Ctx) {
		fl := ctx.GetFriendList()
		fmt.Println(fl.String())
		ctx.SendChain(message.Text("欢迎使用iceinu 冰犬Bot"))
	})
}

func init() {
	// 触发插件注册
	RegisterPlugin(&ExamplePlugin{})
}
