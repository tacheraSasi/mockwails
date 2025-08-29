package config

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

var (
	once sync.Once
)
const (
	AppVersion = "1.0.0"
	AppName    = "MockWails"
	AppID      = "com.tacheraSasi.mockwails"
	Author     = "Tachera Sasi"
)

type Config struct {
	ConfigPath string
	DBPath     string
	DebugMode  bool
	OsType     string
}

type AppDetails struct {
	Name    string
	Version string
	Author  string
}

func init() {
	fmt.Println("Config initialized...")
	once.Do(func() {
		config := GetConfig()
		if err := os.MkdirAll(config.ConfigPath, 0755); err != nil {
			fmt.Printf("Failed to create config directory: %v\n", err)
		}
	})
}

func GetAppDetails() *AppDetails {
	return &AppDetails{
		Name:    AppName,
		Version: AppVersion,
		Author:  Author,
	}
}

func GetConfig() *Config {
	configPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath += "/.mockwails"
	return &Config{
		ConfigPath: configPath,
		DBPath:     configPath + "/db.sqlite",
		DebugMode:  false,
		OsType:     getOsType(),
	}
}

func getOsType() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return "Unknown"
	}
}
