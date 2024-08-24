package plugins

import (
	"fmt"
	"strconv"
	
	"github.com/KyokuKong/go-iceinu/bot/core"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/zerobot"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type RolesPlugin struct{}

// 角色等级映射表
var roleNames = map[int]string{
	0: "普通用户",
	1: "授权用户",
	2: "管理员",
	3: "超级管理员",
}

func (p *RolesPlugin) PluginInfos() (string, string, string, string) {
	var pluginId = "iceinu-roles-manager-plugin"    // 插件id，不能重复
	var pluginVersion = "1.0.1"                     // 插件版本
	var pluginIntroduction = "用于进行权限管理插件" // 插件简介
	var pluginDeveloper = "Kyoku"                   // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *RolesPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	
	// getrole命令实现
	engine.OnRegex(fmt.Sprintf(`^%sgetrole(?:\s+(\S+))?$`, cfg.Bot.CommandPrefix)).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)
		// 如果没有传入参数，显示自己的权限组
		if regexMatched[1] == "" {
			roleLevel, _ := core.GetUserRole(ctx.Event.UserID)
			roleName := roleNames[int(roleLevel)]
			ctx.SendChain(message.At(ctx.Event.UserID), message.Text(fmt.Sprintf(" 你的权限组是: %v", roleName)))
		} else {
			// 传入了参数，显示指定用户的权限组
			if !core.CheckUserRole(ctx.Event.UserID, 2) {
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 你没有权限查看其他用户的权限组，需要权限：管理员"))
			} else {
				qid, _ := strconv.Atoi(regexMatched[1])
				if qid != 0 {
					roleLevel, _ := core.GetUserRole(int64(qid))
					roleName := roleNames[int(roleLevel)]
					ctx.SendChain(message.At(ctx.Event.UserID), message.Text(fmt.Sprintf(" 用户 %v 的权限组是: %v", regexMatched[1], roleName)))
				} else {
					ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" getrole命令的参数类型有误"))
				}
			}
		}
	})
	
	// setrole命令实现
	engine.OnRegex(fmt.Sprintf(`^%ssetrole\s+(\d+)\s+(\d+)$`, cfg.Bot.CommandPrefix)).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)
		targetID, _ := strconv.ParseInt(regexMatched[1], 10, 64)
		newRole, _ := strconv.Atoi(regexMatched[2])
		
		// 获取操作者和目标用户的权限等级
		operatorRoleLevel, _ := core.GetUserRole(ctx.Event.UserID)
		
		// 检查是否尝试对自己进行操作
		if ctx.Event.UserID == targetID {
			ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 你不能对自己的权限组进行修改"))
			return
		}
		
		// 检查权限
		if operatorRoleLevel > int16(newRole) || (operatorRoleLevel == 3 && int16(newRole) == 3) {
			// 操作者权限高于目标权限 或者 操作者和目标用户均为超级管理员
			err := core.SetUserRole(targetID, int16(newRole))
			if err != nil {
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 设置权限失败，请稍后重试"))
			} else {
				roleName := roleNames[newRole]
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(fmt.Sprintf(" 用户 %v 的权限组已设置为: %v", regexMatched[1], roleName)))
			}
		} else {
			ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 你没有足够的权限来设置此用户的权限组"))
		}
	})
}

func (p *RolesPlugin) PluginHelps() []models.CommandHelp {
	return []models.CommandHelp{
		{
			IsShown:     true,
			CommandName: "getrole",
			Usage:       fmt.Sprintf("%sgetrole {QQ号}", cfg.Bot.CommandPrefix),
			Description: "查看自己或指定用户所处的权限组",
			Flags:       map[string]string{},
		},
		{
			IsShown:     true,
			CommandName: "setrole",
			Usage:       fmt.Sprintf("%ssetrole {QQ号} {等级}", cfg.Bot.CommandPrefix),
			Description: "设置指定用户的权限组",
			Flags:       map[string]string{},
		},
	}
}

func init() {
	// 触发插件注册
	core.RegisterPlugin(&RolesPlugin{})
}
