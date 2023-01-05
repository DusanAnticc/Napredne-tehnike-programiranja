package config

import (
	"fmt"

	"github.com/DusanAnticc/Napredne-tehnike-programiranja/UserService/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
    hostdb     = "localhost"
    userdb     = "postgres"
    passworddb = "root"
    dbname  = "NtpDB"
)

var users = []model.User{
	{
		Model:        gorm.Model{},
		Username:     "dusan",
		Password:     "123",
		EmailAddress: "dusan@gmail.com",
		FirstName:    "Dusan",
		LastName:     "Antic",
	},
	{
		Model:          gorm.Model{},
		Username:       "mirko",
		Password:     	"123",
		EmailAddress:   "mirko@gmail.com",
		FirstName:      "Mirko",
		LastName:       "Mirkovic",
	},
	{
		Model:          gorm.Model{},
		Username:       "marko",
		Password:     	"123",
		EmailAddress:   "marko@gmail.com",
		FirstName:      "Marko",
		LastName:       "Markovic",
	},
	{
		Model:          gorm.Model{},
		Username:       "nikola",
		Password:     	"123",
		EmailAddress:   "nikola@gmail.com",
		FirstName:      "Nikola",
		LastName:       "Nikolic",
	},
}


func initDb(){

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", hostdb, userdb, passworddb, dbname)	
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
	} else {
		fmt.Println("Connected to database!")
	}

	db.Migrator().DropTable("users")
	db.Migrator().AutoMigrate(&model.User{})

	for _, user := range users {
		db.Create(&user)
	}
}
