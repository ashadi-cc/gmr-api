package service

import (
	"api-gmr/auth"
	"api-gmr/model"
	"api-gmr/repository"
	"api-gmr/repository/mysql"
	"fmt"
	"log"
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
		userRepo: mysql.NewUserRepo(),
	}
}

func (service *AuthService) Validate(data model.UserLogin) (model.User, error) {
	var user model.User

	dbUser, err := service.userRepo.FindByUsername(data.Username)
	if err != nil {
		log.Println(err.Error())
		return user, fmt.Errorf("could not find username")
	}

	//validate password
	if !auth.CheckPasswordHash(data.Password, dbUser.GetPasswordHash()) {
		return user, fmt.Errorf("invalid password")
	}

	user = model.User{
		Id:    dbUser.GetUserID(),
		Email: dbUser.GetEmail(),
		Group: dbUser.GetUserGroup(),
	}

	return user, nil
}

func (service *AuthService) CreateToken(user model.User) (string, error) {
	tokenString, err := auth.CreateToken(&user)
	if err != nil {
		log.Printf("failed created token with error: %s \n", err.Error())
		return "", fmt.Errorf("internal Server Error")
	}
	return tokenString, nil
}
