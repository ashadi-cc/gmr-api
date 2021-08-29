package auth

import (
	"api-gmr/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	//AppName represents APP Name
	AppName = config.GetApp().AppName
	//ExpiredDuration token expired duration in hours
	ExpiredDuration = time.Duration(24*365) * time.Hour
	//JwtSigningMethod jwt signing method choose
	JwtSigningMethod = jwt.SigningMethodHS256
	//SecretKey jwt secret key value
	SecretKey = config.GetApp().JwtSecret
	//JwtSignatureKey signature key from secret key
	JwtSignatureKey = []byte(SecretKey)
)

//Claim containts common jwt standar claims and user data
type Claim struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Group    string `json:"group"`
}
