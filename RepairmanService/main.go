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
	handler := adapter.NewRepairmanHandler(dataservice)
	router.HandleRequests(handler)
}
