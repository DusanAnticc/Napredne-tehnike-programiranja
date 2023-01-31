package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	roundrobin "github.com/hlts2/round-robin"
)

var BaseUserService, _ = roundrobin.New(&url.URL{Host: "http://localhost:8081"})
var BaseRepairmanService, _ = roundrobin.New(&url.URL{Host: "http://localhost:8082"})
var BaseAdminService, _ = roundrobin.New(&url.URL{Host: "http://localhost:8083"})
var BaseEmailService, _ = roundrobin.New(&url.URL{Host: "http://localhost:8000"})

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func SetupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func AuthorizeAdmin(r *http.Request) bool {

	request, _ := http.NewRequest(http.MethodGet, BaseUserService.Next().Host+"/api/users/authA", nil)
	bytes.NewBufferString("")
	request.Header.Set("Accept", "application/json")
	values := r.Header.Values("Authorization")

	if len(values) == 0 {
		return false
	}

	request.Header.Set("Authorization", values[0])
	authClient := &http.Client{}
	authResponse, err := authClient.Do(request)

	if err != nil {
		return false
	}

	if authResponse.StatusCode != 200 {
		return false
	}

	return true
}

func AuthorizeUser(r *http.Request) bool {

	request, err := http.NewRequest(http.MethodGet, BaseUserService.Next().Host+"/api/users/authS", nil)

	request.Header.Values("Authorization")
	request.Header.Set("Authorization", request.Header.Values("Authorization")[0])

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if response.StatusCode != 200 {
		fmt.Println("Unauthorized")
		return false
	}

	return false
}
