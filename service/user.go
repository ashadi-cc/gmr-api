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

type IUserService interface {
	UserInfo(userID int) (model.User, error)
	UpdateUser(user model.User) error
}

type UserService struct {
	userRepo repository.User
}

func NewUserService() IUserService {
	return &UserService{
		userRepo: repo().GetUserRepository(),
	}
}

func (service *UserService) UserInfo(userID int) (model.User, error) {
	var user model.User

	dbUser, err := service.userRepo.FindByUserID(context.Background(), userID)
	if err != nil {
		if cause := errors.Cause(err); cause == sql.ErrNoRows {
			return user, util.NewUserError(http.StatusBadRequest, "user id not found", err)
		}
		return user, err
	}

	user = model.User{
		Id:       dbUser.GetUserID(),
		Email:    dbUser.GetEmail(),
		Group:    dbUser.GetGroup(),
		Username: dbUser.GetUsername(),
		Blok:     dbUser.GetBlok(),
		Name:     dbUser.GetName(),
	}
	return user, nil
}

func (service *UserService) UpdateUser(user model.User) error {
	if user.Password != "" {
		hashPasword, err := auth.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = hashPasword
	}

	err := service.userRepo.UpdateEmailandPassword(context.Background(), user)
	return err
}
