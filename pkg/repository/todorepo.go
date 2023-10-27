package repository

import (
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) repo.RepoInterfaces {
	return &TodoRepository{collection: collection}

}

func (rr *TodoRepository) CreateName(request models.Test) (string, error) {
	cmd := bson.D{
		{"insert", "Todo"},
		{"documents", []interface{}{request}},
	}
	result := bson.M{}
	err := rr.collection.Database().RunCommand(context.Background(), cmd).Decode(&result)
	if err != nil {
		return "", err
	}
	if result["ok"] == 1 {
		return "Data Inserted Successfully", nil
	} else {
		return "failed to insert data", nil
	}
}
