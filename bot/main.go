package main

import (
	"fmt"
	
	"github.com/KyokuKong/go-iceinu/bot/config"
)

func main() {
	// 程序主函数，在这里运行整个初始化以及启动流程
	cfg, err := config.GetConfig("./config.toml")
	if err != nil {
		return
	}
	fmt.Println(cfg.Bot.AutoUpdate)
}
