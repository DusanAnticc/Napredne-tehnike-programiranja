package model

type UpdateUserDto struct {
	Id           string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	EmailAddress string `gorm:"not null"`
}

type RepairmanDto struct {
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	Username      string `gorm:"not null;unique"`
	EmailAddress  string `gorm:"not null"`
	Role          string
	Price         uint64 `gorm:"not null"`
	Occupation    string `gorm:"not null"`
	Address       string
	Review        float32
	ReviewCounter uint64
}

type CreateRepairmanDto struct {
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	Username      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	EmailAddress  string `gorm:"not null"`
	Price         uint64 `gorm:"not null"`
	Occupation    string `gorm:"not null"`
	Address       string
	Review        float32
	ReviewCounter uint64
}

type UserDto struct {
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Username     string `gorm:"not null;unique"`
	EmailAddress string `gorm:"not null"`
	Role         string
	Address      string
	Banned       bool `gorm:"default:false"`
}
