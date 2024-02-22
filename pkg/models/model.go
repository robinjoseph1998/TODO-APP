package models

type Task struct {
	NO   uint   `bson:"no`
	Task string `bson:"task"`
}

type User struct {
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	Phone     string `bson:"phone"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
}
