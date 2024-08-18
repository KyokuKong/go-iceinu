// Package db 一些用于初始化数据库连接并对数据库进行操作的函数
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/models"
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
		log.Info("当前使用的数据库类型为：SQLite")
		// 自动切片，由于GORM不支持读sqlite开头的URL
		dsn = cfg.Database.SqlConnect[10:]
		log.Debugf("目标数据库URL：%s", dsn)
		dialector = sqlite.Open(dsn)
	case cfg.Database.SqlConnect[:8] == "postgres":
		log.Info("当前使用的数据库类型为：PostgreSQL")
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.Database.SqlUsername, cfg.Database.SqlPassword, cfg.Database.SqlConnect[11:], cfg.Database.SqlDatabase)
		log.Debugf("目标数据库URL：%s", dsn)
		dialector = postgres.Open(dsn)
	case cfg.Database.SqlConnect[:5] == "mysql":
		log.Info("当前使用的数据库类型为：MySQL")
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

// MigrateTables 自动迁移表结构
func MigrateTables() {
	db := DB
	// 检查并创建/更新 users 表
	if !db.Migrator().HasTable(&models.User{}) {
		log.Info("表 'users' 不存在，正在创建...")
		if err := db.Migrator().CreateTable(&models.User{}); err != nil {
			log.Fatalf("创建表 'users' 失败: %v", err)
		}
	} else {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("更新表 'users' 失败: %v", err)
		}
	}

	// 检查并创建/更新 events_log 表
	if !db.Migrator().HasTable(&models.EventLog{}) {
		log.Info("表 'events_log' 不存在，正在创建...")
		if err := db.Migrator().CreateTable(&models.EventLog{}); err != nil {
			log.Fatalf("创建表 'events_log' 失败: %v", err)
		}
	} else {
		if err := db.AutoMigrate(&models.EventLog{}); err != nil {
			log.Fatalf("更新表 'events_log' 失败: %v", err)
		}
	}
}
