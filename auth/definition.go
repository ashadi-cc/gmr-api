package auth

import (
	"api-gmr/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AppName = "API_GMR"
)

var (
	ExpiredDuration  = time.Duration(24*365) * time.Hour
	JwtSigningMethod = jwt.SigningMethodHS256
	SecretKey        = env.GetValue("JWT_SECRET", "abcd")
	JwtSignatureKey  = []byte(SecretKey)
)

type Claim struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Group  string `json:"group"`
}
