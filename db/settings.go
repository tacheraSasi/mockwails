package db

import "time"

type Settings struct {
	ID                     uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AllowDedicatedPorts    bool      `gorm:"not null;default:false" json:"allowDedicatedPorts"`
	DefaultUnifiedPort     int       `gorm:"not null;default:8080" json:"defaultUnifiedPort"`
	CreatedAt              time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Settings) TableName() string {
	return "settings"
}