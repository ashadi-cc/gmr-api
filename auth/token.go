package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

//UserInterface represents user methods interface
type UserInterface interface {
	GetUserID() int
	GetEmail() string
	GetGroup() string
	GetUsername() string
	SetUserId(id int)
	SetEmail(email string)
	SetGroup(group string)
	SetUsername(string)
}

//CreateToken returns string token by given user model
func CreateToken(user UserInterface) (string, error) {
	claim := Claim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    AppName,
			ExpiresAt: time.Now().Add(ExpiredDuration).Unix(),
		},
		UserId:   user.GetUserID(),
		Email:    user.GetEmail(),
		Group:    user.GetGroup(),
		Username: user.GetUsername(),
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claim,
	)

	signedToken, err := token.SignedString(JwtSignatureKey)
	if err != nil {
		return "", errors.Wrap(err, "can't create token")
	}

	return signedToken, nil
}

//ValidateToken returns Jwt Claims by given token
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
		return nil, errors.Wrap(err, "token is invalid")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claim, nil
}

//ClaimToUser set user field values from claim
func ClaimToUser(claim jwt.MapClaims, user UserInterface) {
	if userId, ok := claim["user_id"].(float64); ok {
		user.SetUserId(int(userId))
	}

	if username, ok := claim["username"].(string); ok {
		user.SetUsername(username)
	}

	if userEmail, ok := claim["email"].(string); ok {
		user.SetEmail(userEmail)
	}

	if userGroup, ok := claim["group"].(string); ok {
		user.SetGroup(userGroup)
	}
}
