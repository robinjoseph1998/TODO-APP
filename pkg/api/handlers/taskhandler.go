package handlers

import (
	"Todo/pkg/api/utils"
	"Todo/pkg/models"
	use "Todo/pkg/usecase/interfaces"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	userID := utils.GetUserIDFromContext(c)
	fmt.Println("USERID CONTEXT", userID)
	id, err := primitive.ObjectIDFromHex(userID)
	EnteredTask := models.Task{
		Task:   Task,
		UserID: id,
	}
	fmt.Println("EWNTERED TASK ID ", EnteredTask.UserID)
	CreatedTask, err := hh.usecase.ExecuteCreateTask(EnteredTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Created Task Is": CreatedTask})
}
