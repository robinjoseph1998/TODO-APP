package interfaces

import "Todo/pkg/api/models"

type RepoInterfaces interface {
	CreateName(request models.Test) (string, error)
}
