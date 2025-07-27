package route

import (
	"task_management/Delivery/controllers"
	domain "task_management/Domain"
	infrastructure "task_management/Infrastructure"
	repository "task_management/Repository"
	usecases "task_management/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func TaskRouter(db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controllers.TaskController{
		TaskUseCase: usecases.NewTaskUseCase(tr),
	}

	group.POST("/tasks", infrastructure.RoleMiddleware("Admin"), tc.CreateTask)
	group.GET("/tasks", tc.GetTasks)
	group.GET("/tasks/:id", tc.GetTaskById)
	group.PUT("/tasks/:id", infrastructure.RoleMiddleware("Admin"), tc.ReplaceTask)
	group.DELETE("/tasks/:id", infrastructure.RoleMiddleware("Admin"), tc.DeleteTask)

}
