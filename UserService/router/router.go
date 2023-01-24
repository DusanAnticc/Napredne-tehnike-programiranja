package router

import (
	"UserService/adapter"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *adapter.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/authA", handler.AuthorizeAdmin).Methods(http.MethodGet)
	router.HandleFunc("/api/users/authS", handler.AuthorizeUser).Methods(http.MethodGet)
	router.HandleFunc("/api/users/get-all-users", handler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/users/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/update", handler.UpdateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/CreateRepairman", handler.CreateRepairman).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-repairman", handler.GetAllRepairman).Methods(http.MethodGet)
	router.HandleFunc("/api/users/ban/{id}", handler.BanUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", router))
}
