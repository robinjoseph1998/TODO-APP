package interfaces

import "Todo/pkg/models"

type UserRepoInterfaces interface {
	CreateUser(user models.User) error
	FetchEmail(email string) (*models.User, error)
	FetchPhoneNumber(phone string) (*models.User, error)
}
