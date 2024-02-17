package handlers

import (
	"Todo/pkg/api/middleware"
	"Todo/pkg/api/utils"
	"Todo/pkg/models"
	use "Todo/pkg/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase use.UserUseCaseInterface
}

func NewUserHandler(usecase use.UserUseCaseInterface) *UserHandler {
	return &UserHandler{usecase}
}

func (uh *UserHandler) UserSignup(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input! please check your provided data and try again"})
		return
	}
	NewUser, err := uh.UserUsecase.ExecuteSignup(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup Successfull",
		"user":    NewUser,
	})
}

func (uh *UserHandler) UserLogin(c *gin.Context) {
	var request utils.Login
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in login credentials"})
		return
	}
	phone := request.Phone
	password := request.Password

	userID, err := uh.UserUsecase.ExecuteLogin(phone, password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid phone number or password"})
		return
	} else {
		middleware.GenToken(userID, phone, c)
		c.JSON(http.StatusOK, gin.H{"message": "Logged In Successfully"})
	}
}
