package config

import (
	"RepairmanService/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	hostdb     = "localhost"
	userdb     = "postgres"
	passworddb = "root"
	dbname     = "NtpDB"
)

var appointments = []model.Appointment{
	{
		Model:             gorm.Model{},
		Username:          "dusan",
		Repairman:         "nikola",
		DateOfAppointment: "2023-01-01",
		TimeOfAppointment: "12:00",
		Accepted:          false,
		Price:             100,
	},
	{
		Model:             gorm.Model{},
		Username:          "dusan",
		Repairman:         "marko",
		DateOfAppointment: "2023-01-01",
		TimeOfAppointment: "13:00",
		Accepted:          false,
		Price:             150,
	},
}

func InitDatabase() *gorm.DB {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", hostdb, userdb, passworddb, dbname)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
	} else {
		fmt.Println("Connected to database!")
	}

	db.Migrator().DropTable("appointments")
	db.Migrator().AutoMigrate(&model.Appointment{})

	for _, appointment := range appointments {
		db.Create(&appointment)
	}
	return db
}
