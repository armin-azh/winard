package models

import (
	"time"
)

type AbstractUser struct {
	AbstractModel
	prime     string `gorm:"unique"`
	username  string `gorm:"unique"`
	password  string `gorm:"unique"`
	lastLogin time.Time
	isActive  bool
}
