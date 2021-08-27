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

type IAuthService interface {
	Validate(user model.UserLogin) (model.User, error)
	CreateToken(user model.User) (string, error)
}

type AuthService struct {
	userRepo repository.User
}

func NewAuthService() IAuthService {
	return &AuthService{
		userRepo: repo().GetUserRepository(),
	}
}

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

func (service *AuthService) CreateToken(user model.User) (string, error) {
	tokenString, err := auth.CreateToken(&user)
	return tokenString, err
}
