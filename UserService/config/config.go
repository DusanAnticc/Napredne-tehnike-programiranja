package config

import (
	"UserService/model"
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

var users = []model.User{
	{
		Model:        gorm.Model{},
		Username:     "dusan",
		Password:     "$2a$10$yu08QPMBib4Jd8i6pUVeeuH.OE22cQWvpGGI5Oqc4DQdRvQL.9Rsm",
		EmailAddress: "dusan@gmail.com",
		FirstName:    "Dusan",
		LastName:     "Antic",
		Role:         model.Admin,
	},
	{
		Model:         gorm.Model{},
		Username:      "mirko",
		Password:      "$2a$10$yu08QPMBib4Jd8i6pUVeeuH.OE22cQWvpGGI5Oqc4DQdRvQL.9Rsm",
		EmailAddress:  "mirko@gmail.com",
		FirstName:     "Mirko",
		LastName:      "Mirkovic",
		Role:          model.Repairman,
		Price:         100,
		Occupation:    "Electrical",
		Address:       "Banja Luka",
		Review:        4,
		ReviewCounter: 21,
	},
	{
		Model:         gorm.Model{},
		Username:      "marko",
		Password:      "$2a$10$yu08QPMBib4Jd8i6pUVeeuH.OE22cQWvpGGI5Oqc4DQdRvQL.9Rsm",
		EmailAddress:  "marko@gmail.com",
		FirstName:     "Marko",
		LastName:      "Markovic",
		Role:          model.Repairman,
		Price:         150,
		Occupation:    "Woodworker",
		Address:       "Banja Luka",
		Review:        4.3,
		ReviewCounter: 26,
	},
	{
		Model:        gorm.Model{},
		Username:     "nikola",
		Password:     "$2a$10$yu08QPMBib4Jd8i6pUVeeuH.OE22cQWvpGGI5Oqc4DQdRvQL.9Rsm",
		EmailAddress: "nikola@gmail.com",
		FirstName:    "Nikola",
		LastName:     "Nikolic",
		Role:         model.Standard,
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

	db.Migrator().DropTable("users")
	db.Migrator().AutoMigrate(&model.User{})

	for _, user := range users {
		db.Create(&user)
	}
	return db
}
