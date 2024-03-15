package repository

import (
	database "Todo/pkg/db"
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository() repo.TaskRepoInterfaces {
	database.ConnectDB()
	collection := database.DB.Database("myproject").Collection("Todo")
	return &TaskRepository{collection}
}

func (rr *TaskRepository) GetTasks(userID string) ([]models.Task, error) {
	filter := bson.M{"userid": userID}
	ctx := context.TODO()
	cur, err := rr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var documents []models.Task
	for cur.Next(ctx) {
		var document models.Task
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

func (rr *TaskRepository) CreateTask(enteredTask models.Task) (*models.Task, error) {
	_, err := rr.collection.InsertOne(context.TODO(), enteredTask)
	if err != nil {
		return nil, err
	}
	var CreatedTask models.Task
	err = rr.collection.FindOne(context.TODO(),
		bson.M{"task": enteredTask.Task, "userid": enteredTask.UserID}).Decode(&CreatedTask)
	if err != nil {
		return nil, err
	}
	return &CreatedTask, nil
}

func (rr *TaskRepository) UpdateTask(taskId string, task string) (*models.Task, error) {
	objId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"task": task}}
	_, err = rr.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	var updatedTask models.Task
	err = rr.collection.FindOne(context.TODO(), filter).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func (rr *TaskRepository) DeleteTask(taskId string, userId string) ([]models.Task, error) {
	objId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId, "userid": userId}
	_, err = rr.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var RestTasks []models.Task
	RestTasks, err = rr.GetTasks(userId)
	if err != nil {
		return nil, err
	}
	return RestTasks, nil
}

func (rr *TaskRepository) DeleteAllTasks(userId string) ([]models.Task, error) {
	filter := bson.M{"userid": userId}
	_, err := rr.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var RestTasks []models.Task
	RestTasks, err = rr.GetTasks(userId)
	if err != nil {
		return nil, err
	}
	return RestTasks, nil
}
