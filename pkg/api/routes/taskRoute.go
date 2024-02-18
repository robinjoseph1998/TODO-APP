package routes

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, TaskHandler *handlers.TaskHandler) *gin.Engine {

	r.GET("/test", middleware.ValidateCookie, TaskHandler.TestFunction)
	r.POST("/addname", middleware.ValidateCookie, TaskHandler.AddName)
	r.GET("/showname", middleware.ValidateCookie, TaskHandler.ShowName)
	r.POST("/write", middleware.ValidateCookie, TaskHandler.CreateTask)

	return r
}
