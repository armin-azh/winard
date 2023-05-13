package models

import (
	"gorm.io/gorm"
	"time"
)

type AbstractUser struct {
	gorm.Model
	prime     string `gorm:"unique"`
	username  string `gorm:"unique"`
	password  string `gorm:"unique"`
	lastLogin time.Time
}
