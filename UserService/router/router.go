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
	router.HandleFunc("/api/users/createRepairman", handler.CreateRepairman).Methods(http.MethodPost)
	router.HandleFunc("/api/users/get-all-repairman", handler.GetAllRepairman).Methods(http.MethodGet)
	router.HandleFunc("/api/users/ban/{username}", handler.BanUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/find/{id}", handler.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/users/payment/{username}/{money}", handler.PaymentOnAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/users/payBill/{username}/{money}", handler.PayBill).Methods(http.MethodPost)
	router.HandleFunc("/api/users/assessment/{username}/{grade}", handler.Assessment).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", router))
}
