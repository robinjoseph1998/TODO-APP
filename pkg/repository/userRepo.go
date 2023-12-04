package repository

import (
	database "Todo/pkg/db"
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() repo.TaskRepoInterfaces {
	database.ConnectDB()
	collection := database.DB.Database("myproject").Collection("Todo")
	return &TaskRepository{collection}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	_, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	log.Println("User Added SuccessFully")
	return nil
}
