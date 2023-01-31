package model

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin     = "Admin"
	Standard  = "Standard"
	Repairman = "Repairman"
	Employee  = "Employee"
)

type User struct {
	gorm.Model
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	Username      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	EmailAddress  string `gorm:"not null"`
	Role          Role
	MoneyBalance  uint64
	Price         uint64
	Occupation    string
	Address       string
	Review        float32
	ReviewCounter uint64
	Banned        bool `gorm:"default:false"`
}

type Login struct {
	Username string
	Password string
}
