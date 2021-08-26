package controller

import (
	"api-gmr/model"
	"api-gmr/service"
	"encoding/json"
	"net/http"
)

type Login struct {
	authService service.IAuthService
}

func NewLogin(authService service.IAuthService) *Login {
	return &Login{
		authService: authService,
	}
}

func (l Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userLogin model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.CommonMessage{Message: "invalid payload"})
		return
	}

	if err := userLogin.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := model.CommonMessage{Success: false, Message: "invalid payload"}.WithError(err)
		json.NewEncoder(w).Encode(data)
		return
	}

	//validate user
	user, err := l.authService.Validate(userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := model.CommonMessage{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	//create token
	tokenString, err := l.authService.CreateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := model.CommonMessage{Success: false, Message: err.Error()}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: tokenString})
}
