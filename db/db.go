package db

import (
	"log"
	"sync"

	"github.com/tacheraSasi/mockwails/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		DBPath := config.GetConfig().DBPath
		db, err = gorm.Open(sqlite.Open(DBPath), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
	})
	return db
}
