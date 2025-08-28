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

type Config struct {
	ConfigPath string
	DBPath     string
	DebugMode  bool
	OsType     string
}

func init() {
	fmt.Println("Config initialized...")
	once.Do(func() {
		// TODO: I will Create a folder mockwails in the $HOME directory of the user
		config := GetConfig()
		os.MkdirAll(config.ConfigPath, os.ModePerm)
	})
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
