package handlers

import (
	"ApiGateway/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Post(utils.BaseRepairmanService.Next().Host+"/api/repairman/createAppointment", "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func AcceptAppointment(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	username := params["username"]

	response, err := http.Post(utils.BaseRepairmanService.Next().Host+"/api/repairman/acceptAppointment/"+idStr+"/"+username, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeclineAppointment(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]
	username := params["username"]

	response, err := http.Post(utils.BaseRepairmanService.Next().Host+"/api/repairman/declineAppointment/"+idStr+"/"+username, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func PayAppointment(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	idStr := params["id"]

	response, err := http.Post(utils.BaseRepairmanService.Next().Host+"/api/repairman/payAppointment/"+idStr, "application/json", r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllAppointment(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)

	response, err := http.Get(utils.BaseRepairmanService.Next().Host + "/api/repairman/findAllAppointment/" + params["username"])

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllAppointmentUser(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)

	response, err := http.Get(utils.BaseRepairmanService.Next().Host + "/api/repairman/findAllAppointmentUser/" + params["username"])

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
