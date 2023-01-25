package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ClientUsername    string `gorm:"not null"`
	RepairmanUsername string `gorm:"not null"`
	Text              string `gorm:"not null"`
	ResponseId        string
	Deleted           bool `gorm:"default:false"`
	Report            bool `gorm:"default:false"`
}
