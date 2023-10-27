package server

import (
	"Todo/pkg/api/di"
	route "Todo/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	Handler := di.InitializeHandlerApi()
	router = route.Routers(router, Handler)
	router.Run(":5000")
}
