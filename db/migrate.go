package db

import (
	"log"
)

func AutoMigrate() {
	db := GetDB()
	err := db.AutoMigrate(&Server{}, &AddressAssigned{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
