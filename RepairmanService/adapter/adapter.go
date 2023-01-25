package adapter

import (
	"RepairmanService/dataService"
	"RepairmanService/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RepairmanHandler struct {
	ds *dataService.DataService
}

func NewRepairmanHandler(ds *dataService.DataService) *RepairmanHandler {
	return &RepairmanHandler{ds}
}

func (uh *RepairmanHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var newReview model.Appointment
	json.NewDecoder(r.Body).Decode(&newReview)

	err := uh.ds.CreateAppointment(&newReview)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newReview)
	}
}

func (uh *RepairmanHandler) AcceptAppointment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idStr := params["id"]
	username := params["username"]
	id, _ := strconv.ParseUint(idStr, 10, 64)

	result, err := uh.ds.AcceptAppointment(id, username)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func (rh *RepairmanHandler) FindAllAppointments(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username := params["username"]
	allUsers := rh.ds.FindAllAppointments(username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}
