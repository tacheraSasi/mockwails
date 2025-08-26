package db

import (
	"time"
)

type Server struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	Description     string    `gorm:"type:text;not null" json:"description"`
	Endpoint        string    `gorm:"type:varchar(255);not null" json:"endpoint"`
	Method          string    `gorm:"type:varchar(10);not null" json:"method"`
	RequestHeaders  string    `gorm:"type:text;not null" json:"requestHeaders"`
	RequestBody     string    `gorm:"type:text;not null" json:"requestBody"`
	ResponseStatus  int       `gorm:"type:int;not null" json:"responseStatus"`
	ResponseHeaders string    `gorm:"type:text;not null" json:"responseHeaders"`
	ResponseBody    string    `gorm:"type:text;not null" json:"responseBody"`
	Status          string    `gorm:"type:varchar(10);default:'inactive'" json:"status"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Server) TableName() string {
	return "servers"
}
