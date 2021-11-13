package models

import "time"

type RefreshSession struct {
	RefreshSecret string    `gorm:"column:refresh_secret;primary_key"`
	UserId        int32     `gorm:"column:user_id;index"`
	IP            string    `gorm:"column:ip"`
	UserAgent     string    `gorm:"column:user_agent"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}