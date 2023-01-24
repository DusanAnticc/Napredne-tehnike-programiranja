package router

import (
	"RepairmanService/adapter"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *adapter.UsersHandler) {
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8082", router))
}
