package interfaces

import (
	"Todo/pkg/api/models"
)

type UseCaseInterface interface {
	ExecuteAddName(request models.Test) (string, error)
}
