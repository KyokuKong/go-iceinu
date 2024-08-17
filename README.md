# Go-Iceinu
基于[Zerobot](https://github.com/wdvxdr1123/ZeroBot)框架开发的QQ Bot，暂时还在开发中。

## 简介

使用go语言编写，分发二进制文件，部署简易，依赖无忧，同时享受更快的响应速度，更好的性能表现。

`IceInu`是一个跨群聊的QQ机器人，包含针对私聊和群聊可用的功能和高度可定义的使用体验。

## 部署

在部署`iceinu-bot`前，你需要准备如下环境：
- 一个`PostgreSQL`/`MySQL`数据库， 不过也可以直接使用内置的`SQLite`嵌入式数据库（但是可能会在数据量变大之后遇到显著的I/O瓶颈）
- 一个支持`Onebot v11`协议的客户端实现（如`Lagrange`，`Napcat`等）

然后可以通过CI获取最新的`Go-Iceinu`自动构建程序。

根据数据库部署指南部署数据库并创建一个符合需求的数据库环境。

启动二进制可执行文件并修改自动生成的配置

## 管理

`iceinu-bot`可以直接通过对其发送命令来实现远程管理，同时也提供一套用于精细化远程管理的API，可以通过对API发送请求来管理Bot的运行状态。

项目内也包含一个使用python开发的cli管理工具，封装了这些API的请求和一些其他的测试功能。

## 编译

```shell
git clone git@github.com:KyokuKong/go-iceinu.git
cd go-iceinu-bot
go mod tidy
go build
```

什么，你问我编译怎么就这么简单？

因为go编译就是这么简单啊(笑