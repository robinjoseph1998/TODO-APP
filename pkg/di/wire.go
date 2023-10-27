//go:build wireinject
// +build wireinject

package di

import (
	"Todo/pkg/api/handlers"
	"Todo/pkg/db"
	"Todo/pkg/repository"
	"Todo/pkg/usecase"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeHandlerApi() *handlers.Handler {
	wire.Build(
		db.ConnectDB,
		ProvideMongoCollection,
		handlers.NewHandler,
		repository.NewTodoRepository,
		usecase.NewUsecase,
	)
	return &handlers.Handler{}

}

func ProvideMongoCollection() (*mongo.Collection, error) {
	collection, err := db.CreateMongoCollection("mydb", "Todo")
	if err != nil {
		return nil, err
	}
	return collection, nil
}
