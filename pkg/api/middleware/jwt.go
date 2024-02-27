package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type YourClaims struct {
	userID string `json:"userID"`
	phone  string `json:"phone"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 10

func GenToken(userID string, phone string, c *gin.Context) (string, error) {
	SecretKey := os.Getenv("SecretKey")
	claims := YourClaims{
		userID: userID,
		phone:  phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Creator",
		},
	}
	fmt.Println("claims", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	expiringSeconds := int(TokenExpireDuration.Seconds())
	c.SetCookie("Authorize", tokenString, expiringSeconds, "", "", false, true)
	return tokenString, nil

}

func ValidateCookie(c *gin.Context) {
	tokenstring, err := c.Cookie("Authorize")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Message":    "Unauthorized User",
		})
		return
	}
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("error", err)
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("SecretKey")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"StatusCode": 401,
				"Message":    "Cookie Expired",
			})
			return
		}
		c.Set("userID", fmt.Sprint(claims["userID"]))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Message":    "Invalid Claims",
		})
		return
	}
}

func DeleteCookie(c *gin.Context) error {
	c.SetCookie("Authorize", "", 0, "", "", false, true)
	fmt.Println("Cookie Deleted")
	return nil
}
