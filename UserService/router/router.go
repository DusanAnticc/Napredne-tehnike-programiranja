package router

import (
	"UserService/adapter"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *adapter.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/auth", handler.Authorize).Methods(http.MethodGet)
	router.HandleFunc("/api/users/get-all-users", handler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/users/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/update", handler.UpdateUser).Methods(http.MethodPost)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}