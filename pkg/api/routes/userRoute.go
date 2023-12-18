package routes

import (
	"Todo/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, UserHandler *handlers.UserHandler) *gin.Engine {
	r.POST("/signup", UserHandler.UserSignup)
	r.POST("/login", UserHandler.UserLogin)

	return r
}
