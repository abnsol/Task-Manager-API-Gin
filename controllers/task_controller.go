package controllers

import (
	"net/http"
	"task_management/data"
	"task_management/models"

	"github.com/gin-gonic/gin"
)

// invoke data.GetTasks
func GetTasks(c *gin.Context) {
	tasks := data.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

// invoke data.GetTasksById
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task associated with the given id not found"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

// invoke data.ReplaceTask
func ReplaceTask(c *gin.Context) {
	id := c.Param("id")

	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	task, err := data.ReplaceTask(id, newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task not found"})
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

// invoke data.DeleteTask
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	message, err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": message})
	}
}

// invoke data.CreateTask
func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}

	task, err := data.CreateTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusCreated, task)
	}
}
