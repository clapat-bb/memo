package model

import (
	"fmt"

	"github.com/clapat-bb/memo/config"
	"github.com/clapat-bb/memo/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBname,
		config.Config.Database.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("database connet failed: %v", err)
	}
	logger.Log.Info("database connect success!")
	AutoMigrate()
}
