package route

import (
	infrastructure "task_management/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db mongo.Database, router *gin.Engine) {
	publicRouter := router.Group("")
	// All Public APIs
	SignupRouter(db, publicRouter)
	LoginRouter(db, publicRouter)
	UpdateUserRoleRouter(db, publicRouter)

	protectedRouter := router.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(infrastructure.AuthMiddleware())
	// All Private APIs
	TaskRouter(db, protectedRouter)

	router.Run("localhost:8080")
}
