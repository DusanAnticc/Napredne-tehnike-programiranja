package handlers

import (
	"fmt"
	"net/http"

	utils "ApiGateway/utils"
)


func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println(utils.BaseUserService.Next().Host)
	response, err := http.Post(utils.BaseUserService.Next().Host+"/login", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
