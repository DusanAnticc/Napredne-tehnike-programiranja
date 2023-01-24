package main

import (
	"RepairmanService/adapter"
	"RepairmanService/config"
	"RepairmanService/dataService"
	"RepairmanService/router"
)

func main() {
	db := config.InitDatabase()
	dataservice := dataService.NewDataService(db)
	handler := adapter.NewUsersHandler(dataservice)
	router.HandleRequests(handler)
}
