package handlers

import (
	"Todo/pkg/api/utils"
	"Todo/pkg/models"
	use "Todo/pkg/usecase/interfaces"
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
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	Id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	EnteredTask := models.Task{
		Task:   Task,
		UserID: Id,
	}
	CreatedTask, err := hh.usecase.ExecuteCreateTask(EnteredTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Created Task Is": CreatedTask})
}
