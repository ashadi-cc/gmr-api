package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AppName   = "API_GMR"
	SecretKey = "gmr-0098xdxdd"
)

var (
	ExpiredDuration  = time.Duration(1) * time.Hour
	JwtSigningMethod = jwt.SigningMethodHS256
	JwtSignatureKey  = []byte(SecretKey)
)

type Claim struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Group  string `json:"group"`
}
