package interfaces

import "Todo/pkg/models"

type UserRepoInterfaces interface {
	CreateUser(user models.User) error
	FetchEmail(email string) (bool, error)
	FetchPhoneNumber(phone string) (bool, error)
	FetchUser(phone string) (*models.User, error)
}
