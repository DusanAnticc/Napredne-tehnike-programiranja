package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique"`
	Password       string `gorm:"not null"`
	EmailAddress   string `gorm:"not null"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
}

var users = []User{
	{
		Model:        gorm.Model{},
		Username:     "dusan",
		Password:     "123",
		EmailAddress: "admin@gmail.com",
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

const (
    host     = "localhost"
    user     = "postgres"
    password = "root"
    dbname   = "NtpDB"
)

func main(){

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)	
	//dsn := "host=localhost user=postgres password=root dbname=exampleDipl"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
	} else {
		fmt.Println("Connected to database!")
	}

	db.Migrator().DropTable("users")
	db.Migrator().AutoMigrate(&User{})

	for _, user := range users {
		db.Create(&user)
	}

}
