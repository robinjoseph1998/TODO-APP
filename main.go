package main

import (
	routes "Todo/pkg/api/routes"
	di "Todo/pkg/di"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	TaskHandler := di.InitializeTaskApi()
	UserHandler := di.InitializeUserApi()
	router = routes.TaskRoutes(router, TaskHandler)
	router = routes.UserRoutes(router, UserHandler)
	router.Run(":5000")

}
