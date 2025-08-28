package db

import (
	"log"
)

// MigrateToUnifiedPorts removes the unique constraint on the port field
// to allow multiple servers to share the same port in unified mode
func MigrateToUnifiedPorts() error {
	db := GetDB()
	
	// Drop the unique index on port if it exists
	// Note: SQLite doesn't support dropping constraints directly,
	// so we need to recreate the table
	
	log.Println("Migrating database for unified port support...")
	
	// Step 1: Create a new temporary table without the unique port constraint
	err := db.Exec(`
		CREATE TABLE IF NOT EXISTS address_assigned_new (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			server_id INTEGER NOT NULL UNIQUE,
			port INTEGER NOT NULL,
			FOREIGN KEY (server_id) REFERENCES servers(id) ON UPDATE CASCADE ON DELETE CASCADE
		)
	`).Error
	if err != nil {
		return err
	}
	
	// Step 2: Copy data from old table to new table
	err = db.Exec(`
		INSERT INTO address_assigned_new (id, server_id, port)
		SELECT id, server_id, port FROM address_assigned
	`).Error
	if err != nil {
		return err
	}
	
	// Step 3: Drop old table
	err = db.Exec(`DROP TABLE address_assigned`).Error
	if err != nil {
		return err
	}
	
	// Step 4: Rename new table to original name
	err = db.Exec(`ALTER TABLE address_assigned_new RENAME TO address_assigned`).Error
	if err != nil {
		return err
	}
	
	log.Println("Database migration completed successfully")
	return nil
}