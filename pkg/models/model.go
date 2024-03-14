package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID string             `bson:"userid"`
	Task   string             `bson:"task"`
}

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string        `bson:"firstname"`
	LastName  string        `bson:"lastname"`
	Phone     string        `bson:"phone"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
}

type Login struct {
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}
