package service

import (
	"api-gmr/auth"
	"api-gmr/model"
	"api-gmr/repository"
	"api-gmr/repository/mysql"
	"context"
	"fmt"
	"log"
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
		userRepo: mysql.NewUserRepo(),
	}
}

func (service *UserService) UserInfo(userID int) (model.User, error) {
	var user model.User

	dbUser, err := service.userRepo.FindByUserID(context.Background(), userID)
	if err != nil {
		log.Println(err.Error())
		return user, fmt.Errorf("internal server error")
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
			log.Println(err.Error())
			return fmt.Errorf("internal server error")
		}

		user.Password = hashPasword
	}

	err := service.userRepo.UpdateEmailandPassword(context.Background(), user)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("internal server error")
	}

	return nil
}
