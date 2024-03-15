package interfaces

import (
	"Todo/pkg/models"
)

type TaskUseCaseInterface interface {
	ExecuteShowTasks(userID string) ([]models.Task, error)
	ExecuteCreateTask(enteredTask models.Task) (*models.Task, error)
	ExecuteUpdateTask(taskId string, updatedTask string) (*models.Task, error)
	ExecuteDeleteTask(taskId string, userId string) ([]models.Task, error)
	ExecuteDeleteAllTasks(userId string) ([]models.Task, error)
}
