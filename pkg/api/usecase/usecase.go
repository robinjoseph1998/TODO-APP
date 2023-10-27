package usecase

import (
	"Todo/pkg/api/models"
	repo "Todo/pkg/api/repository/interfaces"
	use "Todo/pkg/api/usecase/interfaces"
	"errors"
)

type usecase struct {
	Repo repo.RepoInterfaces
}

func NewUsecase(Repo repo.RepoInterfaces) use.UseCaseInterface {
	return &usecase{Repo: Repo}
}

func (uu *usecase) ExecuteAddName(request models.Test) (string, error) {
	SavedName, err := uu.Repo.CreateName(request)
	if err != nil {
		return "", errors.New("Can't added name")
	}
	return SavedName, nil
}
