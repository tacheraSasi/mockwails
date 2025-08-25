package goofer

type Server struct {
	ID   uint   `orm:"primaryKey;autoIncrement" validate:"required"`
	Name string `orm:"type:varchar(255);notnull" validate:"required"`
	Host string `orm:"type:varchar(255);notnull" validate:"required"`
	Port uint   `orm:"type:int;notnull" validate:"required"`
	RequestCount uint `orm:"type:int;default:0;notnull" validate:"required"`
	ResponseStructure string `orm:"type:text;notnull" validate:"required"`
	RequestStructure string `orm:"type:text;notnull" validate:"required"`
}

func (Server) TableName() string {
	return "servers"
}
