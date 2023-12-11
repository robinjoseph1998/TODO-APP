//go:build wireinject
// +build wireinject

package di

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/repository"
	"Todo/pkg/usecase"

	"github.com/google/wire"
)

func InitializeTaskApi() *handlers.TaskHandler {
	wire.Build(
		handlers.NewTaskHandler,
		usecase.NewTaskUsecase,
		repository.NewTaskRepository,
	)
	return &handlers.TaskHandler{}

}

func InitializeUserApi() *handlers.UserHandler {
	wire.Build(
		handlers.NewUserHandler,
		usecase.NewUserUsecase,
		repository.NewUserRepository,
	)
	return &handlers.UserHandler{}
}
