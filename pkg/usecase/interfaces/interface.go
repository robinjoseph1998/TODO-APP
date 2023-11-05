package interfaces

import (
	"Todo/pkg/models"
)

type UseCaseInterface interface {
	ExecuteAddName(request models.Test) (string, error)
	ExecuteShowName() ([]models.Test, error)
	ExecuteCreateTask(enteredTask models.Task) (*models.Task, error)
}
