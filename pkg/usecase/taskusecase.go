package usecase

import (
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	use "Todo/pkg/usecase/interfaces"
	"errors"
)

type TaskUsecase struct {
	Repo repo.TaskRepoInterfaces
}

func NewTaskUsecase(Repo repo.TaskRepoInterfaces) use.TaskUseCaseInterface {
	return &TaskUsecase{Repo}
}

func (uu *TaskUsecase) ExecuteAddName(request models.Test) (string, error) {
	SavedName, err := uu.Repo.CreateName(request)
	if err != nil {
		return "", errors.New("can't add name")
	}
	return SavedName, nil
}

func (uu *TaskUsecase) ExecuteShowName() ([]models.Test, error) {
	names, err := uu.Repo.GetName()
	if err != nil {
		return nil, errors.New("can't get name")
	}
	return names, nil
}

func (uu *TaskUsecase) ExecuteCreateTask(enteredTask models.Task) (*models.Task, error) {
	CreatedTask, err := uu.Repo.CreateTask(enteredTask)
	if err != nil {
		return nil, errors.New("issues in creating task")
	}
	return CreatedTask, nil
}
