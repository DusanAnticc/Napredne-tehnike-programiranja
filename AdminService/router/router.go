package router

import (
	"AdminService/adapter"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests(handler *adapter.AdminHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/admin/createReview", handler.CreateReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/createReviewResponse/{id}", handler.CreateReviewResponse).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/reportReview/{id}", handler.ReportReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/deleteReview/{id}", handler.DeleteReview).Methods(http.MethodPost)
	router.HandleFunc("/api/admin/get-all-reviews/{username}", handler.GetAllReviews).Methods(http.MethodGet)
	router.HandleFunc("/api/admin/get-all-reviews", handler.GetAll).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8083", router))
}
