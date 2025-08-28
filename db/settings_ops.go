package db

import (
	"errors"
	"gorm.io/gorm"
)

// GetSettings retrieves the settings from the database (creates default if none exist)
func GetSettings() (*Settings, error) {
	db := GetDB()
	var settings Settings
	
	err := db.First(&settings).Error
	if err != nil {
		// If no settings exist, create default settings
		if errors.Is(err, gorm.ErrRecordNotFound) {
			defaultSettings := Settings{
				AllowDedicatedPorts: false, // Default to unified mode
				DefaultUnifiedPort:  8080,  // Default unified port
			}
			if err := db.Create(&defaultSettings).Error; err != nil {
				return nil, err
			}
			return &defaultSettings, nil
		}
		return nil, err
	}
	
	return &settings, nil
}

// UpdateSettings updates the settings in the database
func UpdateSettings(settings *Settings) error {
	db := GetDB()
	return db.Save(settings).Error
}

// IsUnifiedMode returns true if the system is in unified mode (not allowing dedicated ports)
func IsUnifiedMode() (bool, error) {
	settings, err := GetSettings()
	if err != nil {
		return true, err // Default to unified mode on error
	}
	return !settings.AllowDedicatedPorts, nil
}

// GetUnifiedPort returns the configured unified port
func GetUnifiedPort() (int, error) {
	settings, err := GetSettings()
	if err != nil {
		return 8080, err // Default port on error
	}
	return settings.DefaultUnifiedPort, nil
}