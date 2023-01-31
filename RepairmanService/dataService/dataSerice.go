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

func (repo *DataService) AcceptAppointment(id uint64, username string) (*model.Appointment, error) {
	var app model.Appointment
	result := repo.db.Table("appointments").Where("id = ?", id).First(&app)

	if result.Error != nil {
		return nil, errors.New("Appointment cannot be found!")
	}

	if app.Repairman != username {
		return nil, errors.New("Not your appointment")
	}

	app.Accepted = true

	retValue := repo.db.Table("appointments").Save(&app)

	return &app, retValue.Error
}

func (repo *DataService) DeclineAppointment(id uint64, username string) (*model.Appointment, error) {
	var app model.Appointment
	result := repo.db.Table("appointments").Where("id = ?", id).First(&app)

	if result.Error != nil {
		return nil, errors.New("Appointment cannot be found!")
	}

	if app.Repairman != username {
		return nil, errors.New("Not your appointment")
	}

	app.Accepted = false

	retValue := repo.db.Table("appointments").Save(&app)

	return &app, retValue.Error
}

func (repo *DataService) PayAppointment(id uint64) (*model.Appointment, error) {
	var app model.Appointment
	result := repo.db.Table("appointments").Where("id = ?", id).First(&app)

	if result.Error != nil {
		return nil, errors.New("Appointment cannot be found!")
	}

	app.Paid = true

	retValue := repo.db.Table("appointments").Save(&app)

	return &app, retValue.Error
}

func (repo *DataService) FindAllAppointments(username string) []model.Appointment {
	var allAppointments []model.Appointment

	repo.db.Table("appointments").Where("(deleted_at IS NULL) AND (repairman = ?)", username).Find(&allAppointments)

	return allAppointments
}

func (repo *DataService) FindAllAppointmentsUser(username string) []model.Appointment {
	var allAppointments []model.Appointment

	repo.db.Table("appointments").Where("(accepted = true) AND (paid = false) AND (username = ?)", username).Find(&allAppointments)

	return allAppointments
}
