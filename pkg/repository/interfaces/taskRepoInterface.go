package interfaces

import (
	"Todo/pkg/models"
)

type TaskRepoInterfaces interface {
	CreateName(request models.Test) (string, error)
	GetName() ([]models.Test, error)
	CreateTask(enteredTask models.Task) (*models.Task, error)
}
