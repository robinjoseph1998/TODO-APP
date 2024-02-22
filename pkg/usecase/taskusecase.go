package usecase

import (
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	use "Todo/pkg/usecase/interfaces"
	"errors"
)

type TaskUsecase struct {
	TaskRepo repo.TaskRepoInterfaces
}

func NewTaskUsecase(Repo repo.TaskRepoInterfaces) use.TaskUseCaseInterface {
	return &TaskUsecase{Repo}
}

func (uu *TaskUsecase) ExecuteShowTasks() ([]models.Task, error) {
	names, err := uu.TaskRepo.GetTasks()
	if err != nil {
		return nil, errors.New("can't get name")
	}
	return names, nil
}

func (uu *TaskUsecase) ExecuteCreateTask(enteredTask models.Task) (*models.Task, error) {
	CreatedTask, err := uu.TaskRepo.CreateTask(enteredTask)
	if err != nil {
		return nil, errors.New("issues in creating task")
	}
	return CreatedTask, nil
}
