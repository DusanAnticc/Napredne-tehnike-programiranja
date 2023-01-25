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
	router.HandleFunc("/api/users/createRepairman", handlers.Update).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-repairman", handlers.GetAllRepairman).Methods(http.MethodGet)
	router.HandleFunc("/api/users/createRepairman", handlers.CreateRepairman).Methods(http.MethodPost)
	router.HandleFunc("/api/users/ban/{id}", handlers.BanId).Methods(http.MethodPost)
	router.HandleFunc("/api/users/find/{id}", handlers.FindId).Methods(http.MethodGet)
	router.HandleFunc("/api/users/payment/{id}/{money}", handlers.Payment).Methods(http.MethodPost)
	router.HandleFunc("/api/users/payBill/{id}/{money}", handlers.PayBill).Methods(http.MethodPost)
	router.HandleFunc("/api/users/assessment/{id}/{grade}", handlers.Assessment).Methods(http.MethodPost)

	// Admin Service
	router.HandleFunc("/api/admin/createReview", handlers.CreateReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/createReviewResponse/{id}", handlers.CreateReviewResponse).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/reportReview/{id}", handlers.ReportReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/deleteReview/{id}", handlers.DeleteReview).Methods(http.MethodPost)

	// Repairman Service
	router.HandleFunc("/api/repairman/createAppointment", handlers.CreateAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/acceptAppointment/{id}/{username}", handlers.AcceptAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/findAllAppointment/{username}", handlers.FindAllAppointment).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))

}
