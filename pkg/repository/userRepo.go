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

func NewUserRepository() repo.UserRepoInterfaces {
	database.ConnectDB()
	collection := database.DB.Database("myproject").Collection("User")
	return &UserRepository{collection}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	_, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	log.Println("User Added SuccessFully")
	return nil
}

func (ur *UserRepository) FetchEmail(email string) (bool, error) {
	var result models.User
	err := ur.collection.FindOne(context.TODO(), email).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else {
			log.Println("error occured", err)
			return false, nil
		}
	}
	return true, nil
}

func (ur *UserRepository) FetchPhoneNumber(phone string) (bool, error) {
	var result models.User
	err := ur.collection.FindOne(context.TODO(), phone).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else {
			log.Println("error occured", err)
			return false, nil
		}
	}
	return true, nil
}
