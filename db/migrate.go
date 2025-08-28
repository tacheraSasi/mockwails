package db

import (
	"log"
)

func AutoMigrate() {
	db := GetDB()
	
	// First, run the standard GORM migration
	err := db.AutoMigrate(&Server{}, &AddressAssigned{}, &Settings{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	
	// Check if we need to run the unified ports migration
	// This checks if the port field still has a unique constraint
	var count int64
	err = db.Raw(`
		SELECT COUNT(*) FROM sqlite_master 
		WHERE type='index' AND name='idx_address_assigned_port' AND tbl_name='address_assigned'
	`).Scan(&count).Error
	
	if err == nil && count > 0 {
		log.Println("Found old unique port constraint, running migration...")
		err = MigrateToUnifiedPorts()
		if err != nil {
			log.Fatalf("Unified ports migration failed: %v", err)
		}
	}
}
