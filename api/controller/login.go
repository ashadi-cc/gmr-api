package controller

import (
	"api-gmr/model"
	"api-gmr/service"
	"api-gmr/util"
	"encoding/json"
	"net/http"
)

//Login represents Login controller
type Login struct {
	authService service.IAuthService
}

//NewLogin returns new Login instance
func NewLogin(authService service.IAuthService) *Login {
	return &Login{
		authService: authService,
	}
}

//@Summary Login endpoint
//@Description retrieve jwt token by given username and password
//@Tags login
//@Param data body model.UserLogin true "user payload"
//@Accept json
//@Produce json
//@Success 200 {object} model.CommonMessage
//@Router /login [post]
//Authenticate Authenticate handler
func (l Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userLogin model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	if err := userLogin.Validate(); err != nil {
		userError := util.NewUserError(http.StatusBadRequest, "invalid payload", err)
		util.PrintUserError(w, userError)
		return
	}

	//validate user
	user, err := l.authService.Validate(userLogin)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	//create token
	tokenString, err := l.authService.CreateToken(user)
	if err != nil {
		util.PrintUserError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(model.CommonMessage{Success: true, Data: tokenString})
}
