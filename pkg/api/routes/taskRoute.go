package routes

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, TaskHandler *handlers.TaskHandler) *gin.Engine {

	r.GET("/showtasks", middleware.ValidateCookie, TaskHandler.ShowTasks)
	r.POST("/createtask", middleware.ValidateCookie, TaskHandler.CreateTask)
	r.PATCH("/edittask", middleware.ValidateCookie, TaskHandler.EditTask)
	r.DELETE("/deletetask", middleware.ValidateCookie, TaskHandler.DeleteTask)
	r.DELETE("/deleteall", middleware.ValidateCookie, TaskHandler.DeleteAllTasks)

	return r
}
