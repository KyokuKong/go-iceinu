// Package db 一些用于初始化数据库连接并对数据库进行操作的函数
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/KyokuKong/go-iceinu/bot/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseConnectionPool() {
	cfg, _ := config.GetConfig()
	// 根据数据库URL判断数据库类型
	var dsn string
	var dialector gorm.Dialector

	switch {
	case cfg.Database.SqlConnect[:6] == "sqlite":
		log.Infof("当前使用的数据库类型为：%sSQLite%s", utils.Cyan, utils.ResetColor)
		// 自动切片，由于GORM不支持读sqlite开头的URL
		dsn = cfg.Database.SqlConnect[10:]
		log.Debugf("目标数据库URL：%s", dsn)
		dialector = sqlite.Open(dsn)
	case cfg.Database.SqlConnect[:8] == "postgres":
		log.Infof("当前使用的数据库类型为：%sPostgreSQL%s", utils.LightGreen, utils.ResetColor)
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.Database.SqlUsername, cfg.Database.SqlPassword, cfg.Database.SqlConnect[11:], cfg.Database.SqlDatabase)
		log.Debugf("目标数据库URL：%s", dsn)
		dialector = postgres.Open(dsn)
	case cfg.Database.SqlConnect[:5] == "mysql":
		log.Infof("当前使用的数据库类型为：%sMySQL%s", utils.LightBlue, utils.ResetColor)
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.SqlUsername, cfg.Database.SqlPassword, cfg.Database.SqlConnect[7:], cfg.Database.SqlDatabase)
		log.Debugf("目标数据库URL：%s", dsn)
		dialector = mysql.Open(dsn)
	default:
		log.Fatalf("未知的数据库类型: %v", cfg.Database.SqlConnect)
	}
	// 建立数据库连接
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
	}
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(8)
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(cfg.Database.DbConnPool)
	// 设置连接的最大生命周期
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将连接存入内存
	DB = db
	log.Info("数据库连接成功")
}

// GetDatabaseConnection 从连接池中获取一个连接
func GetDatabaseConnection() (*sql.Conn, error) {
	if DB == nil {
		return nil, fmt.Errorf("数据库还没有初始化，请检查初始化执行顺序！")
	}
	// 获取底层的 *sql.DB 对象
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层数据库连接池失败: %v", err)
	}
	// 从连接池中获取一个连接
	conn, err := sqlDB.Conn(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %v", err)
	}

	return conn, nil
}

// GetORM 获取数据库引擎
func GetORM() *gorm.DB {
	return DB
}

// migrateTable 用于检查指定表是否符合需求并自动创建/迁移的函数
func migrateTable(db *gorm.DB, tableName string, model interface{}) {
	// 检查表是否存在
	if !db.Migrator().HasTable(model) {
		log.Infof("表 '%s' 不存在，正在创建...", tableName)
		if err := db.Migrator().CreateTable(model); err != nil {
			log.Fatalf("创建表 '%s' 失败: %v", tableName, err)
		}
	} else {
		log.Infof("表 '%s' 存在，正在更新...", tableName)
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("更新表 '%s' 失败: %v", tableName, err)
		}
	}
}

// MigrateTables 自动迁移表结构
func MigrateTables() {
	db := DB

	// 为每个表调用通用的迁移函数
	migrateTable(db, "users", &models.User{})
	migrateTable(db, "events_log", &models.EventLog{})
	migrateTable(db, "plugins", &models.Plugins{})
}
