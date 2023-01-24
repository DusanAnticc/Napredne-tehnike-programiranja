package dataService

import (
	"UserService/model"
	"errors"
	"unicode"

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

func (repo *DataService) Update(userDto *model.UpdateUserDto) error {
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

func (repo *DataService) FindAllUsers() []model.User {
	var allUsers []model.User

	repo.db.Table("users").Where("(deleted_at IS NULL)").Find(&allUsers)

	return allUsers
}

func (repo *DataService) FindByEmail(email string) (model.User, error) {
	var user model.User

	result := repo.db.Table("users").Where("email_address = ? AND deleted_at IS NULL", email).First(&user)

	return user, result.Error
}

func (repo *DataService) FindById(id uint64) (model.User, error) {
	var user model.User

	result := repo.db.Table("users").Where("id = ? AND deleted_at IS NULL", id).First(&user)

	return user, result.Error
}

func ValidatePassword(password string) bool {
	letterNum := 0
	var Number, CapitalLetter, SpecialSign bool

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			Number = true
			letterNum++
		case unicode.IsUpper(c):
			CapitalLetter = true
			letterNum++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			SpecialSign = true
		case unicode.IsLetter(c) || c == ' ':
			letterNum++
		default:
		}
	}

	if Number && CapitalLetter && SpecialSign && letterNum > 6 {
		return true
	}

	return false
}

func (repo *DataService) CreateRepairman(userDto *model.CreateRepairmanDto) error {
	var user model.User

	/*result :=*/
	repo.db.Table("users").Where("username = ?", userDto.Username).First(&user)
	//err1 := errors.New("record not found")

	//if result.Error != err1 {
	//	return errors.New("User already exist!")
	//}

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.EmailAddress = userDto.EmailAddress
	user.Password = userDto.Password
	user.Role = model.Repairman
	user.Address = userDto.Address
	user.Occupation = userDto.Occupation
	user.Price = userDto.Price

	retValue := repo.db.Table("users").Save(&user)

	return retValue.Error
}

func (repo *DataService) FindAllRepairman() []model.User {
	var allUsers []model.User

	repo.db.Table("users").Where("(deleted_at IS NULL) AND (role = ?)", model.Role("Repairman")).Find(&allUsers)

	return allUsers
}

func (repo *DataService) BanUser(id uint64) (*model.User, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User cannot be found!")
	}

	user.Banned = true

	retValue := repo.db.Table("users").Save(&user)

	return &user, retValue.Error
}

func (repo *DataService) Payment(id, money uint64) (*model.User, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User cannot be found!")
	}

	user.MoneyBalance = user.MoneyBalance + money

	retValue := repo.db.Table("users").Save(&user)

	return &user, retValue.Error
}

func (repo *DataService) PayBill(id, money uint64) (*model.User, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User cannot be found!")
	}

	if (user.MoneyBalance - money) < 0 {
		return nil, errors.New("User don't have enough money in account")
	}

	user.MoneyBalance = user.MoneyBalance - money

	retValue := repo.db.Table("users").Save(&user)

	return &user, retValue.Error
}

func (repo *DataService) Assesment(username string, grade uint) error {
	var user model.User
	result := repo.db.Table("users").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return errors.New("User cannot be found!")
	}

	newGrade := user.Review*float32(user.ReviewCounter) + float32(grade)
	user.ReviewCounter = user.ReviewCounter + 1
	user.Review = newGrade / float32(user.ReviewCounter)

	retValue := repo.db.Table("users").Save(&user)

	return retValue.Error
}
