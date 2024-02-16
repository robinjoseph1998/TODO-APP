package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type YourClaims struct {
	userID int    `json:"userID"`
	phone  string `json:"phone"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 10

func GenToken(userID int, phone string, c *gin.Context) (string, error) {
	SecretKey := os.Getenv("SecretKey")
	claims := YourClaims{
		userID: userID,
		phone:  phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Creator",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}
	expiringSeconds := int(TokenExpireDuration.Seconds())
	c.SetCookie("Authorize", tokenString, expiringSeconds, "", "", false, true)
	return tokenString, nil

}
