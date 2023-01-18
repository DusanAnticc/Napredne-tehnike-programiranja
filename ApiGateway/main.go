package main

import (
	"ApiGateway/router"
	"fmt"
)

func main() {

	fmt.Println("Api gateway started")

	router.HandleRequests()
}