package routes

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, TaskHandler *handlers.TaskHandler) *gin.Engine {

	r.GET("/test", middleware.ValidateCookie, TaskHandler.TestFunction)
	r.GET("/showtasks", middleware.ValidateCookie, TaskHandler.ShowTasks)
	r.POST("/createtask", middleware.ValidateCookie, TaskHandler.CreateTask)

	return r
}
