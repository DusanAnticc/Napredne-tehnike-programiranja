package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Username       string `gorm:"not null;unique"`
	Password       string `gorm:"not null"`
	EmailAddress   string `gorm:"not null"`
}