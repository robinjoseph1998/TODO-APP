package main

import (
	route "Todo/pkg/api/routes"
	"Todo/pkg/di"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	Handler := di.InitializeHandlerApi()
	router = route.Routers(router, Handler)
	router.Run(":5000")

}
