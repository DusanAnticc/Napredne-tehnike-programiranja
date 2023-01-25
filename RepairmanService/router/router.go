package router

import (
	"RepairmanService/adapter"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *adapter.RepairmanHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/repairman/createAppointment", handler.CreateAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/acceptAppointment/{id}/{username}", handler.AcceptAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/findAllAppointment/{username}", handler.FindAllAppointments).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8082", router))
}
