package models

import (
	"time"
)

type AbstractUser struct {
	AbstractModel
	Prime     string `gorm:"unique"`
	Username  string `gorm:"unique"`
	Password  string `gorm:"unique"`
	LastLogin time.Time
	IsActive  bool
}
