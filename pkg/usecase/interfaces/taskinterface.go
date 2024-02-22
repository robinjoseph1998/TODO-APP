package interfaces

import (
	"Todo/pkg/models"
)

type TaskUseCaseInterface interface {
	ExecuteShowTasks() ([]models.Task, error)
	ExecuteCreateTask(enteredTask models.Task) (*models.Task, error)
}
