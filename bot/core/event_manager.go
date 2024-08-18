package core

import (
	"encoding/json"
	"time"

	"github.com/KyokuKong/go-iceinu/bot/config"
	"github.com/KyokuKong/go-iceinu/bot/db"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var cfg, _ = config.GetConfig()

// 创建事件过程的封装
func newEvent(promoter int64, envi int64, eventType string, eventData interface{}) error {
	// 创建一个新的 UUID
	eventID := uuid.New()

	// 获取当前时间
	recordTime := time.Now()

	// 将 eventData 序列化为 JSON 字符串
	eventJSON, err := json.Marshal(eventData)
	if err != nil {
		log.Errorf("序列化事件数据失败: %v", err)
		return err
	}

	// 构建 EventLog 实例
	eventLog := models.EventLog{
		EventID:     eventID,
		RecordTime:  recordTime,
		Promoter:    promoter,
		Environment: envi,
		Type:        eventType,
		Event:       string(eventJSON),
	}

	DB := db.GetORM()

	// 将 eventLog 插入数据库
	DB.Create(eventLog)

	log.Debugf("%v Event %v Created by %v at %v", eventLog.Type, eventLog.EventID, eventLog.Promoter, eventLog.Environment)
	return nil
}

// SendInitializeEvent 发送初始化事件，这个事件应该在Bot启动时被注册
func SendInitializeEvent() error {
	eventData := models.InitializeEvent{
		Nickname: cfg.Bot.Nickname,
		Status:   1,
	}
	// 新建事件
	err := newEvent(-1, -1, "initialize", eventData)
	if err != nil {
		return err
	}
	return nil
}

// SendCreateUserEvent 发送新建用户事件
func SendCreateUserEvent(promoter int64, envi int64, status int8) error {
	eventData := models.CreateUserEvent{
		Status: status,
	}
	// 新建事件
	err := newEvent(promoter, envi, "create_user", eventData)
	if err != nil {
		return err
	}
	return nil
}

// SendDeleteUserEvent 发送删除用户事件
func SendDeleteUserEvent(promoter int64, envi int64, status int8) error {
	eventData := models.DeleteUserEvent{
		Status: status,
	}
	// 新建事件
	err := newEvent(promoter, envi, "delete_user", eventData)
	if err != nil {
		return err
	}
	return nil
}
