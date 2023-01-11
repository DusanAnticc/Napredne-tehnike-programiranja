package dataService

import (
	"UserService/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DataService struct {
	db *gorm.DB
}

func NewDataService(db *gorm.DB) *DataService {
	return &DataService{db}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (repo *DataService) Save(user model.User) (uint, error) {
	hashedPassword, _ := HashPassword(user.Password)

	user.Password = hashedPassword
	result := repo.db.Create(&user)

	return user.ID, result.Error
}

func (repo *DataService) Update(userDto *model.UpdateUserDto) (error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", userDto.Id).First(&user)

	if result.Error != nil {
		return errors.New("User cannot be found!")
	}

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.EmailAddress = userDto.EmailAddress

	retValue := repo.db.Table("users").Save(&user)

	return retValue.Error
}

func (repo *DataService) FindAllUsers() ([]model.User) {
	var allUsers [] model.User

	repo.db.Table("users").Where("(deleted_at IS NULL)").First(&allUsers)


	return allUsers
}

func (repo *DataService) FindByEmail(email string) (model.User, error) {
	var user model.User

	result := repo.db.Table("users").Where("email_address = ?", email).First(&user)

	return user, result.Error
}
