# Go-Iceinu

基于Zerobot框架开发的QQ Bot，暂时还在开发中。

## 部署

在部署`iceinu-bot`前，你需要准备一个`PostgreSQL`/`MySQL`数据库， 不过也可以直接使用内置的`SQLite`嵌入式数据库（但是可能会在数据量变大之后遇到显著的I/O瓶颈）

## 管理

`iceinu-bot`可以直接通过对其发送命令来实现远程管理，同时也提供一套用于精细化远程管理的API，可以通过对API发送请求来管理Bot的运行状态。

项目内也包含一个使用python开发的cli管理工具，封装了这些API的请求和一些其他的测试功能。

## 编译

```shell
git clone git@github.com:KyokuKong/go-iceinu-bot.git
cd go-iceinu-bot
go mod tidy
go build
```

什么，你问我编译怎么就这么简单？

因为go编译就是这么简单啊(笑