package controllers

import (
	"net/http"
	domain "task_management/Domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase domain.ITaskUseCase
}

// invoke tc.TaskUseCase.GetTasks
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.TaskUseCase.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

// invoke tc.TaskUseCase.GetTasksById
func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUseCase.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task associated with the given id not found"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

// invoke tc.TaskUseCase.ReplaceTask
func (tc *TaskController) ReplaceTask(c *gin.Context) {
	id := c.Param("id")

	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	task, err := tc.TaskUseCase.ReplaceTask(id, newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task not found"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

// invoke tc.TaskUseCase.DeleteTask
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	message, err := tc.TaskUseCase.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": message})
	}
}

// invoke tc.TaskUseCase.CreateTask
func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	task, err := tc.TaskUseCase.CreateTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusCreated, task)
	}
}
