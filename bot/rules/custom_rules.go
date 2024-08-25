package rules

import (
	"github.com/KyokuKong/go-iceinu/bot/core"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func CheckNotBanned(pluginId string) zero.Rule {
	return func(ctx *zero.Ctx) bool {
		// 判断是否在群聊中，在私聊则不判断这个命令是否被ban
		if ctx.Event.PostType == "message" && ctx.Event.DetailType == "group" {
			// 获取禁用这个命令的群组列表
			groups, _ := core.GetPluginBannedGroups(pluginId)
			for _, groupId := range groups {
				if ctx.Event.GroupID == groupId {
					// 相等则说明这个群组禁用了这个插件
					return false
				}
			}
			return true
		} else {
			return true
		}
	}
}
