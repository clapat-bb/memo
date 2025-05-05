package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex" json:"name"`
	Memos []*Memo `gorm:"many2many:memo_tags"`
}
