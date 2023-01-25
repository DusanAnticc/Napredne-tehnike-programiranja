package main

import (
	"AdminService/adapter"
	"AdminService/config"
	"AdminService/dataService"
	"AdminService/router"
)

func main() {
	db := config.InitDatabase()
	dataservice := dataService.NewDataService(db)
	handler := adapter.NewAdminHandler(dataservice)
	router.HandleRequests(handler)
}
