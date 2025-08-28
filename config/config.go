package config

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

type Config struct {
	
}

func init(){
	fmt.Println("Config initialized...")
	once.Do(func(){
		// TODO: I will Create a folder mockwails in the $HOME directory of the user
	})
}

func GetConfig() *Config {
	return &Config{}
}
