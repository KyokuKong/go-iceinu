package core

import "github.com/KyokuKong/go-iceinu/bot/models"

// EcoUse 使用用户指定数量的货币
func EcoUse(currency string, user *models.User, value int) (bool, *models.User) {
	// 通过switch语句来判断不同货币的消耗
	switch currency {
	case "silver":
		if value >= user.Silver {
			user.Silver -= value
			return true, user
		}
	case "gold":
		if value >= user.Gold {
			user.Gold -= value
			return true, user
		}
	case "ticket":
		if value >= user.Ticket {
			user.Ticket -= value
			return true, user
		}
	}
	// 如果执行到这里没有跳出则说明交易失败
	return false, user
}

// EcoAdd 增加用户指定数量的货币
func EcoAdd(currency string, user *models.User, value int) *models.User {
	switch currency {
	case "silver":
		user.Silver += value
		return user
	case "gold":
		user.Gold += value
		return user
	case "ticket":
		user.Ticket += value
		return user
	}
	return user
}
