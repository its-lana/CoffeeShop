package dto

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	ID   int    `json:"id"` //this can be merchant_id, customer_id
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(claims JWTClaims) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	claims.Issuer = os.Getenv("APP_NAME")
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
