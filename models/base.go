package models

import "gorm.io/gorm"

type AbstractModel struct {
	gorm.Model
	Version uint64 `gorm:"default:0"`
}
