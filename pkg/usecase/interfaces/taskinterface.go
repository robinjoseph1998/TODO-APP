package interfaces

import (
	"Todo/pkg/models"
)

type TaskUseCaseInterface interface {
	ExecuteAddName(request models.Test) (string, error)
	ExecuteShowName() ([]models.Test, error)
	ExecuteCreateTask(enteredTask models.Task) (*models.Task, error)
}
