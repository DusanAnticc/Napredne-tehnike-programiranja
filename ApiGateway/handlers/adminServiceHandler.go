package handlers

import (
	"github.com/gorilla/mux"
	"net/http"

	utils "ApiGateway/utils"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Post(utils.BaseAdminService.Next().Host+"/api/admin/createReview", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]

	response, err := http.Post(utils.BaseAdminService.Next().Host+"/api/admin/reportReview/"+idStr, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CreateReviewResponse(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]

	response, err := http.Post(utils.BaseAdminService.Next().Host+"/api/admin/createReviewResponse/"+idStr, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]

	response, err := http.Post(utils.BaseAdminService.Next().Host+"/api/admin/deleteReview/"+idStr, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
