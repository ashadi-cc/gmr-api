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

//IUserService represents a service for user methods
type IUserService interface {
	//UserInfo returns model.User by given userID
	UserInfo(userID int) (model.User, error)
	//UpdateUser method for update users by given user model payload.
	UpdateUser(user model.User) error
}

//UserService impelmenting IUserService
type UserService struct {
	userRepo repository.User
}

//NewUserService return a new UserService instance
func NewUserService() IUserService {
	return &UserService{
		userRepo: repo().GetUserRepository(),
	}
}

//UserInfo impelemnting IUserService.UserINfo
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

//UpdateUser implementing IUserService.UpdateUser
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
