package routes

import (
	"Todo/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine, handler *handlers.Handler) *gin.Engine {

	r.GET("/test", handler.TestFunction)
	r.POST("/write", handler.AddName)

	return r
}
