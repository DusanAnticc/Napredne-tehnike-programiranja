package dataService

import (
	"RepairmanService/model"
	"errors"
	"gorm.io/gorm"
)

type DataService struct {
	db *gorm.DB
}

func NewDataService(db *gorm.DB) *DataService {
	return &DataService{db}
}

func (repo *DataService) CreateAppointment(app *model.Appointment) error {
	var newAppointment model.Appointment

	newAppointment.DateOfAppointment = app.DateOfAppointment
	newAppointment.TimeOfAppointment = app.TimeOfAppointment
	newAppointment.Username = app.Username
	newAppointment.Repairman = app.Repairman
	newAppointment.Price = app.Price

	retValue := repo.db.Table("appointments").Save(&newAppointment)

	return retValue.Error
}

func (repo *DataService) AcceptAppointment(id uint64) error {
	var app model.Appointment
	result := repo.db.Table("users").Where("id = ?", id).First(&app)

	if result.Error != nil {
		return errors.New("User cannot be found!")
	}

	app.Accepted = true

	retValue := repo.db.Table("users").Save(&app)

	return retValue.Error
}
