package router

import (
	"ApiGateway/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {
	router := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"OPTIONS", "HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	// User Service
	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handlers.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-users", handlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/users/update", handlers.Update).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-repairman", handlers.GetAllRepairman).Methods(http.MethodGet)
	router.HandleFunc("/api/users/createRepairman", handlers.CreateRepairman).Methods(http.MethodPost)
	router.HandleFunc("/api/users/ban/{username}", handlers.BanId).Methods(http.MethodPost)
	router.HandleFunc("/api/users/find/{id}", handlers.FindId).Methods(http.MethodGet)
	router.HandleFunc("/api/users/payment/{username}/{money}", handlers.Payment).Methods(http.MethodPost)
	router.HandleFunc("/api/users/payBill/{username}/{money}", handlers.PayBill).Methods(http.MethodPost)
	router.HandleFunc("/api/users/assessment/{username}/{grade}", handlers.Assessment).Methods(http.MethodPost)

	// Admin Service
	router.HandleFunc("/api/admin/createReview", handlers.CreateReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/createReviewResponse/{id}", handlers.CreateReviewResponse).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/reportReview/{id}", handlers.ReportReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/deleteReview/{id}", handlers.DeleteReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/get-all-Reviews/{username}", handlers.GetAllReviews).Methods(http.MethodGet)
	router.HandleFunc("/api/admin/get-all-Reviews", handlers.FindAllReviews).Methods(http.MethodGet)

	// Repairman Service
	router.HandleFunc("/api/repairman/createAppointment", handlers.CreateAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/acceptAppointment/{id}/{username}", handlers.AcceptAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/declineAppointment/{id}/{username}", handlers.DeclineAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/payAppointment/{id}", handlers.PayAppointment).Methods(http.MethodPost)
	router.HandleFunc("/api/repairman/findAllAppointment/{username}", handlers.FindAllAppointment).Methods(http.MethodGet)
	router.HandleFunc("/api/repairman/findAllAppointmentUser/{username}", handlers.FindAllAppointmentUser).Methods(http.MethodGet)

	router.HandleFunc("/api/email/send", handlers.Email).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(router)))

}
