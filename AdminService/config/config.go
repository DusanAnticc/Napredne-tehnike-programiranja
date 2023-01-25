package config

import (
	"AdminService/model"
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

var reviews = []model.Review{
	{
		ClientUsername:    "mirko",
		RepairmanUsername: "marko",
		Text:              "Great job",
		Deleted:           false,
		Report:            false,
	},
	{
		ClientUsername:    "mirko",
		RepairmanUsername: "nikola",
		Text:              "Superb job",
		Deleted:           false,
		Report:            false,
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

	db.Migrator().DropTable("reviews")
	db.Migrator().AutoMigrate(&model.Review{})

	for _, review := range reviews {
		db.Create(&review)
	}
	return db
}
