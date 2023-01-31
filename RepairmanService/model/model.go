package model

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	Username          string `gorm:"not null"`
	Repairman         string `gorm:"not null"`
	DateOfAppointment string `gorm:"not null"`
	TimeOfAppointment string `gorm:"not null"`
	Accepted          bool   `gorm:"default:false"`
	Price             uint   `gorm:"min:0"`
	Paid              bool   `gorm:"default:false"`
}
