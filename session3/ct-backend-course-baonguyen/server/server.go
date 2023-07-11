package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	http.HandleFunc("/api/public/register", register)
	http.HandleFunc("/api/public/login", login)
	http.HandleFunc("/api/private/self", self)

	// http.HandleFunc("/api/public/log/register", LogWrapper(register))
	// http.HandleFunc("/api/public/log/login", LogWrapper(login))
	// http.HandleFunc("/api/private/log/self", LogWrapper(self))

	http.ListenAndServe(":8090", nil)
}

/*
		TODO #2:
		- implement the logic to register a new user (username, password, full_name, address)
	  	- Validate username (not empty and unique)
	  	- Validate password (length should at least 8)
*/
func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	
	reqBody, err := io.ReadAll(r.Body)
	if (err != nil) {
		writeBadRequest(w, err)
		return 
	}
	err = json.Unmarshal(reqBody, &req)
	if (err != nil) {
		writeBadRequest(w, err)
		return
	}

	if err := validateUsername(req.Username); err != nil {
		writeBadRequest(w, err)
		return
	}

	if err := validatePassword(req.Password); err != nil {
		writeBadRequest(w, err)
		return
	}
	userInfo := UserInfo{
		Fullname: req.Fullname,
		Address: req.Address,
		Username: req.Username,
		Password: req.Password,
	}
	if err := userStore.Save(userInfo); err != nil {
		writeBadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func writeBadRequest(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err.Error())	
}

func validatePassword (password string) error {
	if len(password) >= 8 {return nil }
	return fmt.Errorf("Password must be at least 8 characters")
}

func validateUsername (username string) error {
	if len(username) == 0 { return fmt.Errorf("username cannot be empty") }
	_, exist := userStore.data[username]
	if exist { return fmt.Errorf("user exists") }
	return nil 
}

type RegisterRequest struct {
	Address string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"full_name"`
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (username, password)
	  	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	reqBody, err := io.ReadAll(r.Body)
	if (err != nil) { 
		writeBadRequest(w, err)
		return 
	}
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		writeBadRequest(w, err)
		return
	}

	user, err := userStore.Get(req.Username)
	if err != nil {
		writeBadRequest(w, err)
		return
	}

	if user.Password != req.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(user.Username, 24*time.Minute)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	resp := LoginResponse{Token: token}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}

/*
TODO #4:
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	extractUserNameFn := func(authenticationHeader string) (string, error) {

		var name string
		token, err := jwt.Parse(authenticationHeader, func(token *jwt.Token) (interface{}, error) {
			// check token signing method etc
			return []byte("ct-secret-key"), nil
		})
		if err != nil {
			return "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name = fmt.Sprint(claims["sub"])
		}

		if name == "" {
			return "", fmt.Errorf("invalid token payload")
		}
		return name, nil
	}

	username, err := extractUserNameFn(authHeader)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	user, _ := userStore.Get(username)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
// func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
// 	panic("TODO implement me")
// }

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/