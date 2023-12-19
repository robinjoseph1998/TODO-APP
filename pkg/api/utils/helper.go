package utils

type Task struct {
	Task string `bson:"task"`
}

type Login struct {
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}
