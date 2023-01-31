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
	router.HandleFunc("/api/repairman/declineAppointment/{id}/{username}", handler.DeclineAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/findAllAppointment/{username}", handler.FindAllAppointments).Methods(http.MethodGet)
	router.HandleFunc("/api/repairman/findAllAppointmentUser/{username}", handler.FindAllAppointmentsUser).Methods(http.MethodGet)
	router.HandleFunc("/api/repairman/payAppointment/{id}", handler.PayAppointment).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8082", router))
}
