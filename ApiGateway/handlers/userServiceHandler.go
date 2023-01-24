package handlers

import (
	"net/http"

	utils "ApiGateway/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Post(utils.BaseUserService.Next().Host+"/api/users/login", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Post(utils.BaseUserService.Next().Host+"/api/users/register", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	if !utils.AuthorizeAdmin(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := http.Get(utils.BaseUserService.Next().Host)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func Update(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Post(utils.BaseUserService.Next().Host+"/api/users/register", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
