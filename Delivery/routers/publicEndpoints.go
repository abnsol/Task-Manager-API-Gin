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

// get user controller
func user_controller(database mongo.Database) *controllers.UserController {
	jwt_service := infrastructure.JwtService{}
	password_service := infrastructure.PasswordService{}
	ur := repository.NewUserRepository(database, domain.CollectionUser)
	uc := &controllers.UserController{
		UserUseCase: usecases.NewUserUseCase(ur, password_service, jwt_service),
	}

	return uc
}

func SignupRouter(database mongo.Database, group *gin.RouterGroup) {
	uc := user_controller(database)
	group.POST("/register", uc.Register)
}

func LoginRouter(database mongo.Database, group *gin.RouterGroup) {
	uc := user_controller(database)
	group.POST("/login", uc.Login)
}

func UpdateUserRoleRouter(database mongo.Database, group *gin.RouterGroup) {
	uc := user_controller(database)
	group.PATCH("/users/role", uc.UpdateUserRole)

}
