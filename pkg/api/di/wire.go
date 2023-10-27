//go:build wireinject
// +build wireinject

package di

import (
	"Todo/pkg/api/db"
	"Todo/pkg/api/handlers"

	"github.com/google/wire"
)

func IntializeHandlerApi() *handlers.Handler {
	wire.Build{

		db.ConnectDB,
		handlers.NewHandler,
	}
	return &handlers.NewHandler{}

}
