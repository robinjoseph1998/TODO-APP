package interfaces

import (
	"Todo/pkg/models"
)

type TaskRepoInterfaces interface {
	GetTasks(userID string) ([]models.Task, error)
	CreateTask(enteredTask models.Task) (*models.Task, error)
	UpdateTask(taskId string, task string) (*models.Task, error)
}
