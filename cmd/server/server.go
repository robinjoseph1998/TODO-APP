package server

import (
	"Todo/pkg/api/di"
	use "Todo/pkg/api/handlers"
	route "Todo/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	Handler := di.IntializeHandlerApi()
	router = route.Routers(router, &use.Handler{})
	router.Run(":5000")
}
