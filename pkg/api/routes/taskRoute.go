package routes

import (
	"Todo/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, TaskHandler *handlers.TaskHandler) *gin.Engine {

	r.GET("/test", TaskHandler.TestFunction)
	r.POST("/write", TaskHandler.AddName)
	r.GET("/showname", TaskHandler.ShowName)
	r.POST("/write", TaskHandler.CreateTask)

	return r
}
