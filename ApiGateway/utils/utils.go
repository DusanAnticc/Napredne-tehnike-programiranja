package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	roundrobin "github.com/hlts2/round-robin"
)

var BaseUserService, _ = roundrobin.New(&url.URL{Host: "http://localhost:8081"})

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

	request, err := http.NewRequest(http.MethodGet, BaseUserService.Next().Host+"/api/users/authA", nil)

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
