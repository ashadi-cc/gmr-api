package service

import (
	"api-gmr/auth"
	"api-gmr/model"
	"api-gmr/store/repository"
	"api-gmr/util"
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
)

//IAuthService represents a service for Authenticate user
type IAuthService interface {
	//Validate validate user by given username and password from model.User
	Validate(user model.UserLogin) (model.User, error)

	//CreateToken create new string token by given user payload
	CreateToken(user model.User) (string, error)
}

//AuthService implementing IAuthService
type AuthService struct {
	userRepo repository.User
}

//NewAuthService returns new a AuthService instance
func NewAuthService() IAuthService {
	return &AuthService{
		userRepo: repo().GetUserRepository(),
	}
}

//Validate implementing IAuthService.Validate
func (service *AuthService) Validate(data model.UserLogin) (model.User, error) {
	var user model.User

	dbUser, err := service.userRepo.FindByUsername(context.Background(), data.Username)
	if err != nil {
		if cause := errors.Cause(err); cause == sql.ErrNoRows {
			return user, util.NewUserError(http.StatusBadRequest, "username not found", err)
		}

		return user, err
	}

	//validate password
	if !auth.CheckPasswordHash(data.Password, dbUser.GetPasswordHash()) {
		return user, util.NewUserError(http.StatusBadRequest, "invalid password", nil)
	}

	user = model.User{
		Id:    dbUser.GetUserID(),
		Email: dbUser.GetEmail(),
		Group: dbUser.GetGroup(),
	}

	return user, nil
}

//CreateToken impelmenting IAuthService.CreateToken
func (service *AuthService) CreateToken(user model.User) (string, error) {
	tokenString, err := auth.CreateToken(&user)
	return tokenString, err
}
