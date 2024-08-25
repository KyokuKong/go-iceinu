package plugins

import (
	"fmt"

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

type Bank struct {
}

func (p *BankPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 编写插件命令实现
	engine.OnCommand("bank").Handle(func(ctx *zero.Ctx) {
		// 获取用户信息
		user, _ := core.GetUserByQID(ctx.Event.UserID)
		// 获取用户的银币数量
		silver := user.Silver
		// 获取用户的金币数量
		gold := user.Gold
		// 获取用户的门票数量
		ticket := user.Ticket
		// 发送消息
		ctx.SendChain(
			message.Text("Iceinu Bank\n>>>\n"), message.At(ctx.Event.UserID), message.Text(fmt.Sprintf("\n你身上现在有：\n%d银币，%d金币，%d兑换券", silver, gold, ticket)),
		)
	})
}

func (p *BankPlugin) PluginHelps() []models.CommandHelp {
	return []models.CommandHelp{
		{
			IsShown:     true,
			CommandName: "bank",
			Usage:       fmt.Sprintf("%sbank", cfg.Bot.CommandPrefix),
			Description: "查看你身上的货币数量",
			Flags:       map[string]string{},
		},
	}
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&BankPlugin{})
}
