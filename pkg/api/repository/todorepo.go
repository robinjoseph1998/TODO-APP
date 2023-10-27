package repository

import (
	"Todo/pkg/api/models"
	repo "Todo/pkg/api/repository/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) repo.RepoInterfaces {
	return &TodoRepository{collection: collection}

}

func (rr *TodoRepository) CreateName(request models.Test) (string, error) {

}
