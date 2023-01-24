package adapter

import (
	"RepairmanService/dataService"
)

type UsersHandler struct {
	ds *dataService.DataService
}

func NewUsersHandler(ds *dataService.DataService) *UsersHandler {
	return &UsersHandler{ds}
}
