package interfaces

import (
	"Todo/pkg/models"
)

type UseCaseInterface interface {
	ExecuteAddName(request models.Test) (string, error)
}
