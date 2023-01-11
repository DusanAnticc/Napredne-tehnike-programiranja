package main

import (
	"UserService/adapter"
	"UserService/config"
	"UserService/dataService"
	"UserService/router"
)

func main(){
	db := config.InitDatabase()
	dataservice := dataService.NewDataService(db)
	handler := adapter.NewUsersHandler(dataservice)
	router.HandleRequests(handler)
}
