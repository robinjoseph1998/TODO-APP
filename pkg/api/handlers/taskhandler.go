package handlers

import (
	"Todo/pkg/models"
	use "Todo/pkg/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase use.TaskUseCaseInterface
}

func NewTaskHandler(usecase use.TaskUseCaseInterface) *TaskHandler {
	return &TaskHandler{usecase}
}

func (hh *TaskHandler) TestFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{":::::": "WELCOME TO MY TODO APP"})
}

func (hh *TaskHandler) ShowTasks(c *gin.Context) {
	Tasks, err := hh.usecase.ExecuteShowTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Tasks": Tasks})
}

func (hh *TaskHandler) CreateTask(c *gin.Context) {
	Task := c.PostForm("task")
	EnteredTask := models.Task{
		Task: Task,
	}
	CreatedTask, err := hh.usecase.ExecuteCreateTask(EnteredTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Created Task Is": CreatedTask})
}
