package adapter

import (
	"UserService/dataService"
	"UserService/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	ds *dataService.DataService
}

func NewUsersHandler(ds *dataService.DataService) *UsersHandler {
	return &UsersHandler{ds}
}

var jwtKey = []byte("DusanBogOtac")

func (uh *UsersHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// claims, _ := token.Claims.(jwt.MapClaims)

	// if claims["role"] != model.Admin {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// }
}

func (uh *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	userCheck, err := uh.ds.FindByEmail(user.EmailAddress)

	if err != nil {
		fmt.Println(err)
	}

	if userCheck.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The email is taken")
		return
	}

	id, saveErr := uh.ds.Save(user)

	w.Header().Set("Content-Type", "application/json")
	if saveErr != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(saveErr.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(id)
	}
}

func (uh *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo model.Login

	json.NewDecoder(r.Body).Decode(&loginInfo)

	user, err := uh.ds.FindByEmail(loginInfo.EmailAddress)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	
	token := jwt.New(jwt.SigningMethodHS256)	
	claims := make(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.EmailAddress
	claims["role"] = user.Role
	claims["expirationTime"] = time.Now().Add(time.Hour * 24).Unix()

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(jwtKey))

	json.NewEncoder(w).Encode(tokenString)
}


func (rh *UsersHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	allUsers := rh.ds.FindAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}


func (uh *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.UpdateUserDto
	json.NewDecoder(r.Body).Decode(&newUser)

	err := uh.ds.Update(&newUser)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newUser)
	}
}

func (uh *UsersHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else{
		userDTO, err := uh.ds.FindById(uint64(id))
		
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		//poslati dto ne citavog usera
		json.NewEncoder(w).Encode(userDTO)
	}
}