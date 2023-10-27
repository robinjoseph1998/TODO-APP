//go:build wireinject
// +build wireinject

package di

import (
	"Todo/pkg/api/db"
	"Todo/pkg/api/handlers"
	"Todo/pkg/api/repository"
	"Todo/pkg/api/usecase"

	"github.com/google/wire"
)

func IntializeHandlerApi() *handlers.Handler {
	wire.Build(

		handlers.NewHandler,
		db.ConnectDB,
		repository.NewTodoRepository,
		usecase.NewUsecase,
	)
	return &handlers.Handler{}

}
