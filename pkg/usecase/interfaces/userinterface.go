package interfaces

import (
	"Todo/pkg/models"
)

type UserUseCaseInterface interface {
	ExecuteSignup(request models.User) (*models.User, error)
	ExecuteLogin(phone, password string) (string, error)
}
