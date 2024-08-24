package core

import (
	"errors"
	"time"
	
	"github.com/KyokuKong/go-iceinu/bot/db"
	"github.com/KyokuKong/go-iceinu/bot/models"
	"gorm.io/gorm"
)

// CreateUser 创建新用户
func CreateUser(qid int64) (*models.User, error) {
	newUser := &models.User{
		QID:          qid,
		Nickname:     "",
		Level:        1,
		Exp:          0,
		Role:         0,
		Subscription: false,
		Silver:       0,
		Gold:         0,
		Ticket:       0,
		Like:         0,
		RegisterDate: time.Now(),
		SignDate:     time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC),
		Backpack:     "{}",
	}
	result := db.DB.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return newUser, nil
}

// DeleteUser 根据 QID 删除用户
func DeleteUser(qid int64) error {
	result := db.DB.Delete(&models.User{}, "q_id = ?", qid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser 更新用户数据
func UpdateUser(user *models.User) error {
	result := db.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByQID 根据 QID 获取用户数据，如果用户不存在则返回两个空值
func GetUserByQID(qid int64) (*models.User, error) {
	var user models.User
	
	// 搜索用户
	result := db.DB.First(&user, "q_id = ?", qid)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 用户不存在，创建新用户
		user, err := CreateUser(qid)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	
	// 返回找到的用户数据
	return &user, nil
}
