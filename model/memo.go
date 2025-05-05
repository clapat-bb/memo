package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`

	Tags []*Tag `gorm:"many2many:memo_tags" json:"tags"`
}
