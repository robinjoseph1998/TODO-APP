package interfaces

import (
	"Todo/pkg/models"
)

type TaskRepoInterfaces interface {
	GetTasks() ([]models.Task, error)
	CreateTask(enteredTask models.Task) (*models.Task, error)
	UpdateTask(taskNo int, task string, userId string) (*models.Task, error)
}
