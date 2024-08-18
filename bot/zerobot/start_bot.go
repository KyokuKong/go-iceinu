package zerobot

import (
	"fmt"

	"github.com/KyokuKong/go-iceinu/bot/config"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func LaunchQQBot() {
	cfg, _ := config.GetConfig()
	// 创建一个空的driver列表
	var drivers []zero.Driver

	// 遍历cfg.Onebot.WebsocketConnects列表，并为每个URL创建一个WebSocketServer driver
	for _, wsURL := range cfg.Onebot.WebsocketConnects {
		d := driver.NewWebSocketServer(16, wsURL, "")
		drivers = append(drivers, d)
	}
	zero.RunAndBlock(&zero.Config{
		NickName:      []string{fmt.Sprint(cfg.Bot.Nickname)},
		CommandPrefix: "/",
		SuperUsers:    cfg.Bot.Superuser,
		Driver:        drivers,
	}, nil)
}
