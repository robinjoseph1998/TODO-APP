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

func (uu *TaskUsecase) ExecuteShowTasks(userID string) ([]models.Task, error) {
	tasks, err := uu.TaskRepo.GetTasks(userID)
	if err != nil {
		return nil, errors.New("can't get tasks")
	}
	return tasks, nil
}

func (uu *TaskUsecase) ExecuteCreateTask(enteredTask models.Task) (*models.Task, error) {
	CreatedTask, err := uu.TaskRepo.CreateTask(enteredTask)
	if err != nil {
		return nil, errors.New("issues in creating task")
	}
	return CreatedTask, nil
}

func (uu *TaskUsecase) ExecuteUpdateTask(taskId string, updatedTask string) (*models.Task, error) {
	Task, err := uu.TaskRepo.UpdateTask(taskId, updatedTask)
	if err != nil {
		return nil, errors.New("can't update task")
	}
	return Task, nil
}
