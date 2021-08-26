package controller

import (
	"api-gmr/model"
	"encoding/json"
	"net/http"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

func (l Login) PostLogin(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
