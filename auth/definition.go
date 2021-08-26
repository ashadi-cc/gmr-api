package auth

import (
	"api-gmr/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	AppName          = config.GetApp().AppName
	ExpiredDuration  = time.Duration(24*365) * time.Hour
	JwtSigningMethod = jwt.SigningMethodHS256
	SecretKey        = config.GetApp().JwtSecret
	JwtSignatureKey  = []byte(SecretKey)
)

type Claim struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Group  string `json:"group"`
}
