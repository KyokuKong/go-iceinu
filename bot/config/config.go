// Package config 一系列用于处理配置文件的函数及一个配置单例，可以实现智能的配置文件读取、生成和自动更新
package config

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	
	"github.com/pelletier/go-toml/v2"
)

// Config 代表配置文件的结构
type Config struct {
	Onebot struct {
		WebsocketConnects []string `toml:"websocket_connects"`
		AccessToken       string   `toml:"access_token"`
	} `toml:"onebot"`
	Database struct {
		SqlConnect string `toml:"sql_connect"`
		DbConnPool int    `toml:"db_conn_pool"`
	} `toml:"database"`
	Bot struct {
		Nickname        string  `toml:"nickname"`
		CommandPrefix   string  `toml:"command_prefix"`
		Superuser       []int64 `toml:"superuser"`
		AutoUpdate      bool    `toml:"auto_update"`
		LogLevel        string  `toml:"log_level"`
		EnableRemoteApi bool    `toml:"enable_remote_api"`
	} `toml:"bot"`
}

var defaultConfig = Config{
	Onebot: struct {
		WebsocketConnects []string `toml:"websocket_connects"`
		AccessToken       string   `toml:"access_token"`
	}{
		WebsocketConnects: []string{"ws://127.0.0.1:23333"},
		AccessToken:       "",
	},
	Database: struct {
		SqlConnect string `toml:"sql_connect"`
		DbConnPool int    `toml:"db_conn_pool"`
	}{
		SqlConnect: "sqlite:///iceinu.db",
		DbConnPool: 10,
	},
	Bot: struct {
		Nickname        string  `toml:"nickname"`
		CommandPrefix   string  `toml:"command_prefix"`
		Superuser       []int64 `toml:"superuser"`
		AutoUpdate      bool    `toml:"auto_update"`
		LogLevel        string  `toml:"log_level"`
		EnableRemoteApi bool    `toml:"enable_remote_api"`
	}{
		Nickname:        "IceInuBot",
		CommandPrefix:   "/",
		Superuser:       []int64{1234567890},
		AutoUpdate:      false,
		LogLevel:        "INFO",
		EnableRemoteApi: true,
	},
}

var (
	instance   *Config
	configPath string
	once       sync.Once
	mu         sync.Mutex // 互斥锁，防止线程之间产生冲突
)

// GetConfig 获取单例配置实例
func GetConfig(path string) (*Config, error) {
	var err error
	once.Do(func() {
		configPath = path
		err = checkAndGenerateConfig(configPath)
		if err == nil {
			instance, err = loadConfig(configPath)
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// SaveConfig 将当前配置写回文件
func SaveConfig() error {
	mu.Lock()
	defer mu.Unlock()
	
	if instance == nil {
		return fmt.Errorf("配置尚未初始化")
	}
	
	configData, err := toml.Marshal(instance)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}
	
	err = os.WriteFile(configPath, configData, 0644)
	if err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}
	
	fmt.Println("配置文件已更新")
	return nil
}

// checkAndGenerateConfig 检测并生成配置文件
func checkAndGenerateConfig(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 如果文件不存在，则创建文件
		fmt.Println("配置文件不存在，正在生成...")
		configData, err := toml.Marshal(defaultConfig)
		if err != nil {
			return fmt.Errorf("生成配置文件失败: %v", err)
		}
		err = os.WriteFile(path, configData, 0644)
		if err != nil {
			return fmt.Errorf("写入配置文件失败: %v", err)
		}
		fmt.Println("配置文件已生成")
	} else {
		// 如果文件存在，则检测并补充缺失的配置项
		fmt.Println("配置文件已存在，正在检查并更新...")
		config, err := loadConfig(path)
		if err != nil {
			return err
		}
		updatedConfig := mergeConfig(defaultConfig, *config)
		configData, err := toml.Marshal(updatedConfig)
		if err != nil {
			return fmt.Errorf("更新配置文件失败: %v", err)
		}
		err = os.WriteFile(path, configData, 0644)
		if err != nil {
			return fmt.Errorf("写入配置文件失败: %v", err)
		}
		fmt.Println("配置文件已更新")
	}
	return nil
}

// loadConfig 读取配置文件
func loadConfig(path string) (*Config, error) {
	var config Config
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}
	err = toml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return &config, nil
}

// mergeConfig 合并配置文件，补充缺失的配置项
func mergeConfig(defaultConfig, userConfig Config) Config {
	mergedConfig := userConfig
	
	// 使用反射检查并补充缺失的配置项
	reflectDefaultConfig := reflect.ValueOf(defaultConfig)
	reflectUserConfig := reflect.ValueOf(userConfig)
	reflectMergedConfig := reflect.ValueOf(&mergedConfig).Elem()
	
	for i := 0; i < reflectDefaultConfig.NumField(); i++ {
		defaultField := reflectDefaultConfig.Field(i)
		userField := reflectUserConfig.Field(i)
		mergedField := reflectMergedConfig.Field(i)
		
		if userField.IsZero() {
			mergedField.Set(defaultField)
		}
	}
	
	return mergedConfig
}
