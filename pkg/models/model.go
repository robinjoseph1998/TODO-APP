package models

type Test struct {
	Name string `bson:"name"`
}

type Task struct {
	ID   uint   `gorm:"PrimaryKey"`
	Task string `bson:"task"`
}
