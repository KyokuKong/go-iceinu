package core

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func LaunchQQBot() {
	zero.RunAndBlock(&zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{123456},
		Driver: []zero.Driver{
			// 正向 WS
			driver.NewWebSocketClient("ws://127.0.0.1:6700", ""),
			// 反向 WS
			driver.NewWebSocketServer(16, "ws://127.0.0.1:6701", ""),
		},
	}, nil)
}
