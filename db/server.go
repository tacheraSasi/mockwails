package db

import (
	"time"
)

type Server struct {
	ID              uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string          `gorm:"type:varchar(255);not null" json:"name"`
	Description     string          `gorm:"type:text;not null" json:"description"`
	Endpoint        string          `gorm:"type:varchar(255);not null" json:"endpoint"`
	Method          string          `gorm:"type:varchar(10);not null" json:"method"`
	RequestHeaders  string          `gorm:"type:text;not null" json:"requestHeaders"`
	RequestBody     string          `gorm:"type:text;not null" json:"requestBody"`
	ResponseStatus  int             `gorm:"type:int;not null" json:"responseStatus"`
	ResponseHeaders string          `gorm:"type:text;not null" json:"responseHeaders"`
	ResponseBody    string          `gorm:"type:text;not null" json:"responseBody"`
	RequestQuery    string          `gorm:"type:text;not null;default:''" json:"requestQuery"`
	Latency         uint            `gorm:"type:int;not null;default:0" json:"latency"`
	Status          string          `gorm:"type:varchar(10);default:'inactive'" json:"status"`

	// One-to-one relation
	AddressAssigned AddressAssigned `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"addressAssigned"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Server) TableName() string {
	return "servers"
}

type AddressAssigned struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	ServerID uint    `gorm:"uniqueIndex;not null" json:"serverId"` // one-to-one with Server
	Port     int     `gorm:"uniqueIndex;not null" json:"port"`     // unique port assigned
	Server   *Server `gorm:"foreignKey:ServerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (AddressAssigned) TableName() string {
	return "address_assigned"
}
