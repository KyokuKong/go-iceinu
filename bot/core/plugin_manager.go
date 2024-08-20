package core

import (
	"errors"
	"sort"

	"github.com/KyokuKong/go-iceinu/bot/db"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Plugin 定义插件接口，插件需要继承这个接口
type Plugin interface {
	PluginInfos() (string, string, string, string)
	PluginCommands()
	PluginHelps() []models.CommandHelp
}

// 存储注册过的插件
var pluginRegistry = make(map[string]Plugin)

// 存储全局帮助信息列表
var helpList []models.CommandHelp

// CheckPluginInDB 在数据库中查找对应的插件是否存在，不存在则注册这个插件,存在则返回对应的插件启用状态
func CheckPluginInDB(pluginName string) (bool, error) {
	var plugin models.Plugins

	// 搜索用户
	result := db.DB.First(&plugin, "plugin_id = ?", pluginName)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 插件没有被注册到数据库中，那么添加这个插件
		log.Infof("检测到了新插件 %s ，将被添加进插件数据库", pluginName)
		newPlugin := &models.Plugins{
			PluginId:  pluginName,
			IsEnabled: true,
		}
		result := db.DB.Create(newPlugin)
		// 将新生成的plugin信息返回
		plugin = *newPlugin
		if result.Error != nil {
			return false, result.Error
		}
	} else if result.Error != nil {
		return false, result.Error
	}
	// 没有异常则返回得到的数据
	return plugin.IsEnabled, nil
}

// RegisterPlugin 插件注册函数，插件需要在init()函数中主动调用这个插件来被注册
func RegisterPlugin(plugin Plugin) {
	id, _, _, _ := plugin.PluginInfos()
	pluginRegistry[id] = plugin
}

// RegisterManager 使用这个函数后会注册当前的插件管理器
func RegisterManager() {
	for id, plugin := range pluginRegistry {
		enabled, err := CheckPluginInDB(id)
		if err != nil {
			log.Warnf("检查插件 %s%s%s 的启用状态时出错: %v\n ，跳过了这个插件的注册", utils.LightCyan, id, utils.ResetColor, err)
			continue
		}
		if enabled {
			log.Infof("插件 %s%s%s 已启用，正在注册命令...\n", utils.LightCyan, id, utils.ResetColor)
			plugin.PluginCommands()

			// 获取插件的帮助信息并添加到全局列表中
			helps := plugin.PluginHelps()
			for _, help := range helps {
				if help.IsShown {
					helpList = append(helpList, help)
				}
			}
		} else {
			log.Infof("插件 %s%s%s 未启用，跳过命令注册。\n", utils.LightCyan, id, utils.ResetColor)
		}
	}

	// 将帮助信息列表按命令名称排序
	sort.Slice(helpList, func(i, j int) bool {
		return helpList[i].CommandName < helpList[j].CommandName
	})
}

// GetHelpList 获取全局帮助信息列表
func GetHelpList() []models.CommandHelp {
	return helpList
}

// PrintHelpList 打印全局帮助信息列表，用于调试
func PrintHelpList() {
	for _, help := range helpList {
		log.Infof("命令: %s\n用法: %s\n功能: %s\nFlags: %v\n",
			help.CommandName, help.Usage, help.Description, help.Flags)
	}
}
