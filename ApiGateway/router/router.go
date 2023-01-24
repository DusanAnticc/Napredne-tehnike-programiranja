package router

import (
	"ApiGateway/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()

	// User Service
	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handlers.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-users", handlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/users/update", handlers.Update).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))

}
