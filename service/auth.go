package service

import (
	"api-gmr/auth"
	"api-gmr/model"
	"fmt"
	"log"
)

type AuthService interface {
	Validate(user model.UserLogin) (model.User, error)
	CreateToken(user model.User) (string, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) Validate(data model.UserLogin) (model.User, error) {
	var user model.User

	//TODO: with db repository
	tmpPassword, _ := auth.HashPassword("12345")
	if !auth.CheckPasswordHash(data.Password, tmpPassword) {
		return user, fmt.Errorf("invalid username/password")
	}
	user = model.User{Id: 1, Email: "test@gmail.com", Group: "user"}
	//END TODO

	return user, nil
}

func (service *authService) CreateToken(user model.User) (string, error) {
	tokenString, err := auth.CreateToken(&user)
	if err != nil {
		log.Printf("failed created token with error: %s \n", err.Error())
		return "", fmt.Errorf("internal Server Error")
	}
	return tokenString, nil
}
