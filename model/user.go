package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"usernaem"`
	Password string `gorm:"not null" json:"password"`
}

func AutoMigrate() {
	DB.AutoMigrate(&User{}, &Memo{}, &Tag{})
}
