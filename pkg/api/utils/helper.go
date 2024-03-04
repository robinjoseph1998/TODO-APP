package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (string, error) {
	idStr := c.GetString("userID")
	if idStr != "" {
		return idStr, nil
	}
	return "", errors.New("user id not found in the context")
}
