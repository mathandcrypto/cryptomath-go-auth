package models

import (
	"time"
)

type RefreshSession struct {
	RefreshSecret string    `gorm:"column:refresh_secret;primary_key;not_null"`
	UserId        int32     `gorm:"column:user_id;index;not_null"`
	IP            string    `gorm:"column:ip;not_null"`
	UserAgent     string    `gorm:"column:user_agent;not_null"`
	CreatedAt     time.Time `gorm:"column:created_at;not_null"`
}

func (RefreshSession) TableName() string {
	return "refresh_sessions"
}