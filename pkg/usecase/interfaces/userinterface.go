package interfaces

import (
	"Todo/pkg/models"
)

type UserUseCaseInterface interface {
	ExecuteSignup(request models.User) (error, models.User)
}
