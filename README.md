# Go-Iceinu 

![Static Badge](https://img.shields.io/badge/Go-1.22.5+-blue?logo=go)
![Static Badge](https://img.shields.io/badge/Python-3.11%2B-green?logo=python)
![workflow](https://github.com/KyokuKong/go-iceinu/actions/workflows/go.yml/badge.svg) 
[![CodeTime Badge](https://img.shields.io/endpoint?style=flat&color=222&url=https%3A%2F%2Fapi.codetime.dev%2Fshield%3Fid%3D25986%26project%3Dgo-iceinu-bot%26in=0)](https://codetime.dev)

基于[Zerobot](https://github.com/wdvxdr1123/ZeroBot)框架开发的QQ Bot，暂时还在开发中。

## 简介

使用go语言编写，分发二进制文件，部署简易，依赖无忧，同时享受更快的响应速度，更好的性能表现。

`IceInu`是一个跨群聊的QQ机器人，包含针对私聊和群聊可用的功能和高度可定义的使用体验。

## 部署

在部署`iceinu-bot`前，你需要准备如下环境：
- 一个`PostgreSQL`数据库， 不过也可以直接使用内置的`SQLite`嵌入式数据库（但是可能会在数据量变大之后遇到显著的I/O瓶颈）
- 一个支持`Onebot v11`协议的客户端实现（如`Lagrange`，`Napcat`等）

然后可以通过CI获取最新的`Go-Iceinu`自动构建。

根据[数据库部署指南](./数据库部署指南.md)部署数据库并创建一个符合需求的数据库环境。

启动二进制可执行文件并修改自动生成的配置文件

Done

## 管理

`Go-Iceinu`提供了一套完整的、RESTful的管理API，可以通过其进行鉴权并远程管理bot的各项功能。

你也可以直接使用[Iceinu Manager](./cli/README.md)，一个使用Python编写的命令行管理工具，封装了管理API的各项功能。

## 编译

```shell
git clone git@github.com:KyokuKong/go-iceinu.git
cd go-iceinu-bot
go mod tidy
go build
```

什么，你问我编译怎么就这么简单？

因为go编译就是这么简单啊(笑

## 进度表

`v1.0.2`
- [x] 内置插件系统的基本设计
- [ ] 权限管理
- [ ] 基本的经济系统
- [x] 优化日志输出表现
- [ ] 准备开始设计API和管理工具
- [ ] 部分自动化测试

`v1.0.1`
- [x] 配置文件生成/读取
- [x] 日志输出
- [x] 数据库连接/初始化
- [x] 数据库处理/提取/交互
- [x] ZeroBot初始化及启动
- [x] CI配置、自动构建

## 数据库结构

### users
存储用户数据的表，可能会随着更新添加新的列

| 列名            | 类型       | 备注                                       |
|---------------|----------|------------------------------------------|
| qid           | bigint   | 主键，用户的QQ号                                |
| nickname      | text     | 用户昵称                                     |
| level         | integer  | 用户等级                                     |
| exp           | integer  | 用户升级所需的经验值                               |
| role          | smallint | 用户的权限组，0为普通用户，1为授权用户，2为管理员，3为超级管理员       |
| subscription  | bool     | 是否启用Bot的推送订阅功能，启用订阅会使Bot主动向用户发送一些通知，默认关闭 |
| silver        | integer  | 用户的普通货币数量                                |
| gold          | integer  | 用户的高级货币数量                                |
| ticket        | integer  | 用户的兑换券数量                                 |
| like          | integer  | 用户发布的内容收到的赞的总数                           |
| register_date | date     | 用户的注册时间（第一次使用bot）                        |
| sign_date     | date     | 用户上一次签到的时间（新用户默认为1971年1月1日）              |
| backpack      | json     | 用户的背包内容，以json键值对的形式存储                    |

### events_log
存储事件日志的表，用于系统记录数据的变更，同时方便查询和溯源

| 列名          | 类型        | 备注                                   |
|-------------|-----------|--------------------------------------|
| event_id    | uuid      | 主键，事件的随机36位UUID                      |
| record_time | timestamp | 事件的记录时间                              |
| promoter    | bigint    | 事件的发起人，存储其qid                        |
| enviroment  | bigint    | 事件的发生环境，如在群聊则为对应的群号，如在私聊则为0          |
| type        | text      | 事件类型，方便快速分类检索                        |
| event       | json      | 使用json类型进行序列化存储的事件内容，每个类型的事件有对应的固定格式 |

## 配置文件示例

```toml
# 这是IceinuBot的参考配置文件
# 仓库地址：https://github.com/KyokuKong/go-iceinu
# 你可以在仓库的README.md中找到和这个一样的配置文件
# ----------------------------------------------------

[onebot] # 与Onebot客户端进行连接时所需的配置
# ----------------------------------------------------
# 配置对onebot的反向ws连接，传入的URL字符串格式为"ws://onebot客户端IP:端口"
# 可以传入多个URL来实现多Bot连接，每个连接都会被IceinuBot注册成一个Driver
websocket_connects = ["ws://127.0.0.1:23333"]
# 配置与onebot客户端进行通信时使用的反向token，推荐使用一个随机的16-64位十六进制字符串
# 留空可以不设置，但是这样可能导致一些难以预料的安全问题
# 同时连接多个Onebot客户端时需要将每个客户端的token都设置成相同的
access_token = ""

[database] # 数据库连接设置
# ----------------------------------------------------
# 配置数据库远程连接的URL
# Iceinu的数据库连接及处理都通过gorm进行，理论上gorm兼容的SQL数据库都可以使用（比如MySQL）
# 但是保证可用性的数据库只有PostgreSQL以及内置的SQLite
sql_connect = "sqlite:///iceinu.db"
# 数据库名称，一般不需要改动
sql_database = "iceinu"
# 数据库用户名，一般不需要改动
sql_username = "iceinubot"
# 数据库密码，一般不需要改动
sql_password = "iceinupassword"
# 连接池上限，设置更多的连接池可以同时向数据库发出更多查询请求，但是会影响性能表现
# 一般不需要设置的太高
db_conn_pool = 10

[bot] # IceinuBot自身的设置
# ----------------------------------------------------
# 设置Bot展示的自身名称
nickname = "IceInuBot"
# 设置命令前缀
command_prefix = "/"
# 设置超级管理员用户，拥有Bot的所有权限，包括设置其他管理员
# Bot会将部分使用日志和消息自动发送到这个账号上
superuser = [1234567890]
# 启用自动更新，会在每天凌晨4点自动从Github上检查是否有新的Release版本并尝试下载
# 如果下载成功那么会尝试自行覆盖并重启Bot
auto_update = false
# Bot向命令行输出日志的等级阈值，一般使用INFO即可
# 当Bot使用量非常巨大以至于输出到命令行的日志会一定程度上影响性能时，可以尝试将其设置到更高的等级
# 但是日志事实上对性能的影响微乎其微，出现这种情况更多可能说明设备需要换新了
log_level = "INFO"
# 是否启用远程管理api，关闭远程管理api后将无法使用iceinu manager
# 远程管理api的默认鉴权用户名为admin，密码为iceinubot，请在启动后立刻修改
# 远程管理api中的每一个账户的权限都和superuser相同，请谨慎将远程管理api的用户给予他人
enable_remote_api = true

```

## API文档

详见[API文档](./API文档.md)

## 更新日志

`v1.0.1`

- 完成 配置文件管理、读取、生成
- 完成 日志输出
- 完成 数据库连接、初始化
- 完成 数据库处理、提取、交互
- 完成 ZeroBot初始化、启动
- 完成 CI配置、自动构建
- 项目基础部分基本完工，可以准备开始写插件功能了