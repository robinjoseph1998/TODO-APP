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
	role   string `json:"role"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 10

func GenToken(userID string, phone string, c *gin.Context) (string, error) {
	SecretKey := os.Getenv("SecretKey")
	claims := YourClaims{
		userID: userID,
		phone:  phone,
		role:   "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Creator",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	expiringSeconds := int(TokenExpireDuration.Seconds())
	c.SetCookie("Authorize", tokenString, expiringSeconds, "", "", false, true)
	fmt.Println("Claims", claims)
	return tokenString, nil
}

func ValidateCookie(c *gin.Context) {
	// Extract JWT token from cookie
	tokenString, err := c.Cookie("Authorize")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Message":    "Unauthorized User",
		})
		return
	}
	//parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("error", err)
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("SecretKey")), nil
	})

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Message":    "Invalid Claims",
		})
	}
	//Check token expiration
	exp, ok := claims["exp"].(float64)
	if !ok || exp < float64(time.Now().Unix()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"Message":    "Token has expired",
		})
		return
	}

	// Extract user ID from claims
	userID, ok := claims["userID"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": 401,
			"Message":    "userid not found in claims",
		})
	}
	// Set user ID in Gin context
	c.Set("userID", userID)
	// Continue processing the request
	c.Next()
}

func DeleteCookie(c *gin.Context) error {
	c.SetCookie("Authorize", "", 0, "", "", false, true)
	fmt.Println("Cookie Deleted")
	return nil
}
