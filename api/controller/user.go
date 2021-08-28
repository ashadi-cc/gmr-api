package controller

import (
	"api-gmr/api/middleware"
	"api-gmr/model"
	"api-gmr/service"
	"api-gmr/util"
	"encoding/json"
	"fmt"
	"net/http"
)

//User represents User Controller
type User struct {
	userService service.IUserService
}

//NewUser returns new User instance
func NewUser(userService service.IUserService) *User {
	return &User{
		userService: userService,
	}
}

//Info User info handler method
func (u User) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	userInfo, err := u.userService.UserInfo(userCtx.Id)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: userInfo})

}

//Update User update hander method
func (u User) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	var userInput model.UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	if err := userInput.Validate(); err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	userCtx.Email = userInput.Email
	userCtx.Password = userInput.Password

	if err = u.userService.UpdateUser(userCtx); err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Message: "user updated"})
}

//Billing user billing handler
func (u User) Billing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userCtx, ok := r.Context().Value(middleware.UserKey).(model.User)
	if !ok {
		util.PrintUserError(w, fmt.Errorf("can't load user from context"))
		return
	}

	data, err := u.userService.GetBilling(userCtx)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: data})
}
