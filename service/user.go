package service

import (
	"api-gmr/model"
	"api-gmr/repository"
	"api-gmr/repository/mysql"
	"context"
	"fmt"
	"log"
)

type IUserService interface {
	UserInfo(userID int) (model.User, error)
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
		log.Println("failed when execute query", err.Error())
		return user, fmt.Errorf("internal server error")
	}

	user = model.User{
		Id:       dbUser.GetUserID(),
		Email:    dbUser.GetEmail(),
		Group:    dbUser.GetUserGroup(),
		Username: dbUser.GetUsername(),
		Blok:     dbUser.GetBlok(),
		Name:     dbUser.GetName(),
	}
	return user, nil
}
