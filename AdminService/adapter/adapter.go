package adapter

import (
	"AdminService/dataService"
	"AdminService/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	ds *dataService.DataService
}

func NewAdminHandler(ds *dataService.DataService) *AdminHandler {
	return &AdminHandler{ds}
}

func (uh *AdminHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var newReview model.Review
	json.NewDecoder(r.Body).Decode(&newReview)

	err := uh.ds.CreateReview(&newReview)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newReview)
	}
}

func (uh *AdminHandler) CreateReviewResponse(w http.ResponseWriter, r *http.Request) {
	var newReview model.Review
	json.NewDecoder(r.Body).Decode(&newReview)

	params := mux.Vars(r)
	idStr := params["id"]

	err := uh.ds.CreateReviewResponse(&newReview, idStr)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newReview)
	}
}

func (uh *AdminHandler) ReportReview(w http.ResponseWriter, r *http.Request) {
	var newReview model.Review
	json.NewDecoder(r.Body).Decode(&newReview)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err := uh.ds.ReportReview(id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newReview)
	}
}

func (uh *AdminHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	var newReview model.Review
	json.NewDecoder(r.Body).Decode(&newReview)

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err := uh.ds.DeleteReview(id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newReview)
	}
}
