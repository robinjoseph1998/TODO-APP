//go:build wireinject
// +build wireinject

package di

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/repository"
	"Todo/pkg/usecase"

	"github.com/google/wire"
)

func InitializeHandlerApi() *handlers.Handler {
	wire.Build(
		handlers.NewHandler,
		usecase.NewUsecase,
		repository.NewTodoRepository,
	)
	return &handlers.Handler{}

}
