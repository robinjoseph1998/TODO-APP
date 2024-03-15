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

type TaskUpdateRequest struct {
	TaskId string `json:"taskId"`
	Task   string `json:"task"`
}

type TaskDeleteRequest struct {
	TaskId string `json:"taskId"`
}
