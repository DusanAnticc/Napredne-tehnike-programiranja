package adapter

import (
	"UserService/dataService"
	"UserService/model"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UsersHandler struct {
	ds *dataService.DataService
}

func NewUsersHandler(ds *dataService.DataService) *UsersHandler {
	return &UsersHandler{ds}
}

var jwtKey = []byte("your-256-bit-secret")

func (uh *UsersHandler) AuthorizeAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie := r.Header.Values("Authorization")

	tokenString := strings.Split(cookie[0], " ")[1]

	claims := make(jwt.MapClaims)
	claims["id"] = ""
	claims["email"] = ""
	claims["role"] = ""
	claims["expirationTime"] = ""
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims, _ = token.Claims.(jwt.MapClaims)

	if claims["role"] != model.Admin {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) AuthorizeUser(w http.ResponseWriter, r *http.Request) {
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

	claims, _ := token.Claims.(jwt.MapClaims)

	if claims["role"] != model.Standard {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (uh *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	userCheck, _ := uh.ds.FindByEmail(user.EmailAddress)

	if userCheck.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The email is taken")
		return
	}

	id, saveErr := uh.ds.Save(user)

	w.Header().Set("Content-Type", "application/json")
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(saveErr.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(id)
	}
}

type LoginResponse struct {
	Token string `json:"Token"`
}

func (uh *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo model.Login
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(r.Body).Decode(&loginInfo)

	user, err := uh.ds.FindByUsername(loginInfo.Username)

	if err == nil {

		claims := make(jwt.MapClaims)
		claims["id"] = user.ID
		claims["email"] = user.EmailAddress
		claims["role"] = user.Role
		claims["expirationTime"] = time.Now().Add(time.Hour * 24).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
		tokenString, _ := token.SignedString(jwtKey)

		fmt.Println(tokenString)

		json.NewEncoder(w).Encode(LoginResponse{Token: tokenString})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
	}
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
	} else {
		user, err := uh.ds.FindById(uint64(id))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		userDto := user.ToUserDTO()

		//poslati dto ne citavog usera
		json.NewEncoder(w).Encode(userDto)
	}
}

func (uh *UsersHandler) CreateRepairman(w http.ResponseWriter, r *http.Request) {
	var newUser model.CreateRepairmanDto
	json.NewDecoder(r.Body).Decode(&newUser)

	err := uh.ds.CreateRepairman(&newUser)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newUser)
	}
}

func (rh *UsersHandler) GetAllRepairman(w http.ResponseWriter, r *http.Request) {

	allUsers := rh.ds.FindAllRepairman()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}

func (uh *UsersHandler) BanUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	user, err := uh.ds.BanUser(username)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (uh *UsersHandler) PaymentOnAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	moneyStr := params["money"]
	money, _ := strconv.ParseUint(moneyStr, 10, 64)
	user, err := uh.ds.Payment(username, money)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (uh *UsersHandler) PayBill(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	moneyStr := params["money"]

	money, _ := strconv.ParseUint(moneyStr, 10, 64)
	user, err := uh.ds.PayBill(username, money)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (uh *UsersHandler) Assessment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameStr := params["username"]
	gradeStr := params["grade"]

	grade, _ := strconv.ParseUint(gradeStr, 10, 64)
	user, err := uh.ds.Assessment(usernameStr, grade)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
