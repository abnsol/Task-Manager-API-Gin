package router

import (
	"task_management/controllers"
	"task_management/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()

	// User Requests
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.PATCH("/users/role", middleware.AuthMiddleware(), middleware.RoleMiddleware("Admin"), controllers.UpdateUserRole)

	// Secure Task Requests
	taskRoutes := router.Group("/tasks", middleware.AuthMiddleware())
	{
		taskRoutes.POST("", middleware.RoleMiddleware("Admin"), controllers.CreateTask)
		taskRoutes.GET("", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskById)
		taskRoutes.PUT("/:id", middleware.RoleMiddleware("Admin"), controllers.ReplaceTask)
		taskRoutes.DELETE("/:id", middleware.RoleMiddleware("Admin"), controllers.DeleteTask)
	}

	router.Run("localhost:8080")
}
