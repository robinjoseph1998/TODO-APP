package main

import (
	route "Todo/pkg/api/routes"
	"Todo/pkg/di"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	TaskHandler := di.InitializeHandlerApi()
	UserHandler := di.InitializeHandlerApi()
	router = route.TaskRoutes(router, TaskHandler)
	router = route.UserRoutes(router, UserHandler)
	router.Run(":5000")

}
