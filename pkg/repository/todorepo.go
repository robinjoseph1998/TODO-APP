package repository

import (
	database "Todo/pkg/db"
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository() repo.RepoInterfaces {
	database.ConnectDB()
	collection := database.DB.Database("myproject").Collection("Todo")
	return &TodoRepository{collection}
}

func (rr *TodoRepository) CreateName(request models.Test) (string, error) {
	_, err := rr.collection.InsertOne(context.TODO(), request)
	if err != nil {
		return "", err
	}
	return "Data Inserted Successfully", nil
}

func (rr *TodoRepository) GetName() ([]models.Test, error) {
	filter := bson.M{}
	ctx := context.TODO()
	cur, err := rr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var documents []models.Test
	for cur.Next(ctx) {
		var document models.Test
		if err := cur.Decode(&document); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}

func (rr *TodoRepository) CreateTask(enteredTask models.Task) (*models.Task, error) {
	fmt.Println("TASK", enteredTask)
	_, err := rr.collection.InsertOne(context.TODO(), enteredTask)
	if err != nil {
		return nil, err
	}
	var CreatedTask models.Task
	fmt.Println("TASK ID", enteredTask.ID)
	err = rr.collection.FindOne(context.TODO(), bson.M{"task": enteredTask.Task}).Decode(&CreatedTask)
	fmt.Println("CREATED TASK", CreatedTask)
	if err != nil {
		return nil, err
	}
	return &CreatedTask, nil
}
