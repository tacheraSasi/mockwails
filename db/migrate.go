package db

import (
	"log"
)

func AutoMigrate() {
	db := GetDB()
	err := db.AutoMigrate(&Server{}, &AddressAssigned{}, &Settings{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
