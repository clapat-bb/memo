package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *AppConfig

type AppConfig struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBname   string
	}
	Jwt struct {
		Secret  string
		Expires int
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config failed: %w", err))
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("parse config failed: %w", err))
	}
}
