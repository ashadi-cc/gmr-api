package controller

import (
	"api-gmr/api/middleware"
	"api-gmr/model"
	"api-gmr/service"
	"encoding/json"
	"net/http"
)

type User struct {
	userService service.IUserService
}

func NewUser(userService service.IUserService) *User {
	return &User{
		userService: userService,
	}
}

func (u User) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		data := model.CommonMessage{Success: false, Message: "internal server error"}
		json.NewEncoder(w).Encode(data)
		return
	}

	userInfo, err := u.userService.UserInfo(userCtx.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data := model.CommonMessage{Success: false, Message: "internal server error"}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: userInfo})

}

func (u User) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		data := model.CommonMessage{Success: false, Message: "internal server error"}
		json.NewEncoder(w).Encode(data)
		return
	}

	var userInput model.UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.CommonMessage{Message: "invalid payload"})
		return
	}

	if err := userInput.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data := model.CommonMessage{Success: false, Message: "invalid payload"}.WithError(err)
		json.NewEncoder(w).Encode(data)
		return
	}

	userCtx.Email = userInput.Email
	userCtx.Password = userInput.Password

	if err = u.userService.UpdateUser(userCtx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data := model.CommonMessage{Success: false, Message: "internal server error"}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Message: "user updated"})
}
