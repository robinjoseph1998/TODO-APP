package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	NO   uint   `bson:"no`
	Task string `bson:"task"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstname"`
	LastName  string             `bson:"lastname"`
	Phone     string             `bson:"phone"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
}
