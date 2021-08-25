package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserInterface interface {
	GetUserId() int
	GetEmail() string
	GetGroup() string
	SetUserId(id int)
	SetEmail(email string)
	SetGroup(group string)
}

func CreateToken(user UserInterface) (string, error) {
	claim := Claim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    AppName,
			ExpiresAt: time.Now().Add(ExpiredDuration).Unix(),
		},
		UserId: user.GetUserId(),
		Email:  user.GetEmail(),
		Group:  user.GetGroup(),
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claim,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("signing method invalid %v", method)
		}
		if method != JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid %v", method)
		}

		return JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claim, nil
}

func ClaimToUser(claim jwt.MapClaims, user UserInterface) {
	if userId, ok := claim["user_id"].(float64); ok {
		user.SetUserId(int(userId))
	}

	if userEmail, ok := claim["email"].(string); ok {
		user.SetEmail(userEmail)
	}

	if userGroup, ok := claim["group"].(string); ok {
		user.SetGroup(userGroup)
	}
}
