package goofer

import "time"

type ServerEntity struct {
	ID              uint   `orm:"primaryKey;autoIncrement" validate:"required"`
	Name            string `orm:"type:varchar(255);notnull" validate:"required"`
	Description     string `orm:"type:text;notnull" validate:"required"`
	Endpoint        string `orm:"type:varchar(255);notnull" validate:"required"`
	Method          string `orm:"type:varchar(10);notnull" validate:"required"`
	RequestHeaders  string `orm:"type:text;notnull" validate:"required"`
	RequestBody     string `orm:"type:text;notnull" validate:"required"`
	ResponseStatus  int    `orm:"type:int;notnull" validate:"required"`
	ResponseHeaders string `orm:"type:text;notnull" validate:"required"`
	ResponseBody    string `orm:"type:text;notnull" validate:"required"`
	Status          string `orm:"type:varchar(10);default:'inactive'"`
	CreatedAt      time.Time `orm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `orm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (ServerEntity) TableName() string {
	return "servers"
}
