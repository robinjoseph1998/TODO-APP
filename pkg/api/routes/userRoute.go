package routes

import (
	"Todo/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, UserHandler *handlers.UserHandler) *gin.Engine {
	r.POST("/signup", UserHandler.UserSignup)

	return r
}
