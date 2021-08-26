package service

import (
	"api-gmr/auth"
	"api-gmr/model"
	"api-gmr/repository"
	"api-gmr/repository/mysql"
	"context"
	"database/sql"
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

	dbUser, err := service.userRepo.FindByUsername(context.Background(), data.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("could not find username")
		}

		log.Println(err)
		return user, fmt.Errorf("internal server error")
	}

	//validate password
	if !auth.CheckPasswordHash(data.Password, dbUser.GetPasswordHash()) {
		return user, fmt.Errorf("invalid password")
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
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("internal Server Error")
	}
	return tokenString, nil
}
