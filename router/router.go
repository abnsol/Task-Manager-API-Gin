package router

import (
	"task_management/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.PUT("/tasks/:id", controllers.ReplaceTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run("localhost:8080")
}
