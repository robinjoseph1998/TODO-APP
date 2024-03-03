package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) string {
	idStr := c.GetString("userID")
	return idStr

}
