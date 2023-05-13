package models

import "gorm.io/gorm"

type AbstractModel struct {
	gorm.Model
	version uint64 `gorm:"default:0"`
}
