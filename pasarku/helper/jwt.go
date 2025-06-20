package helper

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type JWTClaim struct {
	UserID uint
	Role   string
	jwt.StandardClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
