package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	QID          int64     `gorm:"primaryKey"`
	Nickname     string    `gorm:"type:text"`
	Level        int       `gorm:"type:int"`
	Exp          int       `gorm:"type:int"`
	Role         int16     `gorm:"type:smallint"`
	Subscription bool      `gorm:"type:boolean;default:false"`
	Silver       int       `gorm:"type:int"`
	Gold         int       `gorm:"type:int"`
	Ticket       int       `gorm:"type:int"`
	Like         int       `gorm:"type:int"`
	RegisterDate time.Time `gorm:"type:date"`
	SignDate     time.Time `gorm:"type:date;default:'1971-01-01'"`
	Backpack     string    `gorm:"type:json"`
}

type EventLog struct {
	EventID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	RecordTime  time.Time `gorm:"type:timestamp"`
	Promoter    int64     `gorm:"type:bigint"`
	Environment int64     `gorm:"type:bigint"`
	Type        string    `gorm:"type:text"`
	Event       string    `gorm:"type:json"`
}

type Plugins struct {
	PluginId       string `gorm:"type:text;primarykey"`
	IsEnabled      bool   `gorm:"type:boolean;default:true"`
	BannedInGroups string `gorm:"type:json"`
}
