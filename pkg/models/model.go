package models

type Test struct {
	Name string `bson:"name"`
}

type Task struct {
	ID   uint   `gorm:"PrimaryKey"`
	Task string `bson:"task"`
}

type User struct {
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	Phone     string `bson:"phone"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
}
