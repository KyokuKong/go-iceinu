package plugins

import (
	"fmt"
	"strconv"

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
	var pluginVersion = "1.0.0"            // 插件版本
	var pluginIntroduction = "用于展示帮助信息的插件" // 插件简介
	var pluginDeveloper = "Kyoku"          // 插件开发者
	return pluginId, pluginVersion, pluginIntroduction, pluginDeveloper
}

func (p *HelpPlugin) PluginCommands() {
	// 调用Bot引擎
	engine := zerobot.GetBotEngine()
	// 动态生成正则表达式，支持命令前缀
	engine.OnRegex(fmt.Sprintf(`^%s(help|帮助)(?:\s+(\S+))?$`, cfg.Bot.CommandPrefix)).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)

		// 获取帮助信息列表
		helpList := core.GetHelpList()
		totalItems := len(helpList)
		itemsPerPage := 10
		totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage

		// 如果没有传入参数，显示第一页的帮助信息
		if regexMatched[2] == "" {
			page := 1
			paginatedHelp := getPaginatedHelp(helpList, page, itemsPerPage, totalPages)
			ctx.SendChain(message.Text(paginatedHelp))
		} else {
			// 尝试转换为整数，转换成功则是页数
			if num, err := strconv.Atoi(regexMatched[2]); err == nil {
				if num > 0 && num <= totalPages {
					paginatedHelp := getPaginatedHelp(helpList, num, itemsPerPage, totalPages)
					ctx.SendChain(message.Text(paginatedHelp))
				} else {
					paginatedHelp := getPaginatedHelp(helpList, totalPages, itemsPerPage, totalPages)
					ctx.SendChain(message.Text(paginatedHelp))
				}
			} else {
				// 传入的是字符串，需要查询help注册列表中有没有相同的命令
				command := regexMatched[2]
				helpText := getHelpForCommand(helpList, command)
				if helpText != "" {
					ctx.SendChain(message.Text(helpText))
				} else {
					ctx.SendChain(message.Text(fmt.Sprintf("未找到命令: %s\n", command)))
				}
			}
		}
	})
}

// getPaginatedHelp 返回指定页码的帮助信息
func getPaginatedHelp(helpList []models.CommandHelp, page int, itemsPerPage int, totalPages int) string {
	startIndex := (page - 1) * itemsPerPage
	endIndex := startIndex + itemsPerPage
	if endIndex > len(helpList) {
		endIndex = len(helpList)
	}

	var helpText = fmt.Sprintf("Iceinu Help \n(第 %d/%d 页)\n<>为必填参数，{}为可选参数\n>>>\n", page, totalPages)
	for _, help := range helpList[startIndex:endIndex] {
		helpText += fmt.Sprintf("%s   %s\n", help.Usage, help.Description)
	}
	if page < totalPages {
		helpText += fmt.Sprintf("\n输入 '%shelp %d' 查看下一页\n", cfg.Bot.CommandPrefix, page+1)
	}
	return helpText
}

// getHelpForCommand 根据命令名称查找并返回帮助信息，并在有 flag 列表时显示 flag 信息
func getHelpForCommand(helpList []models.CommandHelp, command string) string {
	for _, help := range helpList {
		if help.CommandName == command {
			helpText := fmt.Sprintf("%s>>>\n用法: %s\n描述: %s\n", help.CommandName, help.Usage, help.Description)

			// 检查是否有可用的 flag 列表
			if len(help.Flags) > 0 {
				helpText += "可用参数:\n"
				for flag, description := range help.Flags {
					helpText += fmt.Sprintf("%s: %s\n", flag, description)
				}
			}
			return helpText
		}
	}
	return ""
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
