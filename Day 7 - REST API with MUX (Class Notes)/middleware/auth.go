package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secrateKey = []byte("helloWorld")

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func NewClaims(
	ID int,
	Username string,
	IsAdmin bool, exirationDate time.Time) *Claims {
	return &Claims{
		ID:       ID,
		Username: Username,
		IsAdmin:  IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exirationDate.Unix()},
	}
}

func (c *Claims) Signing() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, _ := token.SignedString(secrateKey)
	return tokenString
}
