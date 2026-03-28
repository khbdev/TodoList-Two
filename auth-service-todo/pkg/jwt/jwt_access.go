package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



var jwtSecret = []byte(os.Getenv("JWT_KEY"))

func GenerateAccesJwtAcccToken(userId uint) (string, error) {
	clams := jwt.MapClaims{
		"user_id": userId, 
		"exp": time.Now().Add(30 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)

	return  token.SignedString(jwtSecret)
}