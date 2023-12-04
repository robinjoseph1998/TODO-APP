package handlers

import (
	"Todo/pkg/models"
	use "Todo/pkg/usecase/interfaces"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase use.UserUseCaseInterface
}

func NewUserHandler(usecase use.UserUseCaseInterface) *UserHandler {
	return &UserHandler{usecase}
}

func (uh *UserHandler) UserSignup(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input! please check your provided data and try again"})
		log.Printf(err.Error())
		return
	}
	NewUser, err := uh.UserUsecase.ExecuteSignup(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Oops! Something went wrong on our end. Please try again"})
		log.Printf(err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup Successfull",
		"user":    NewUser,
	})
}
