package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// collections
const (
	CollectionTask = "Tasks"
	CollectionUser = "Users"
)

// Task Entity
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Time        string `json:"time"`
	Status      bool   `json:"status"`
}

// User Entity
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
	Role     string             `json:"role"`
}

// Task use case interfaces
type ITaskUseCase interface {
	CreateTask(task Task) (string, error)
	GetTasks() []Task
	GetTaskById(id string) (Task, error)
	ReplaceTask(id string, newTask Task) (Task, error)
	DeleteTask(id string) (string, error)
}

// Task repository interfaces
type ITaskRepository interface {
	CreateTask(task Task) error
	GetTasks() []Task
	GetTaskById(id string) (Task, error)
	ReplaceTask(id string, newTask Task) (Task, error)
	DeleteTask(id string) error
}

// User use case interfaces
type IUserUseCase interface {
	Register(user User) (string, error)
	Login(user User) (string, string, error)
	UpdateUserRole(email string, newRole string) error
}

// User repository interfaces
type IUserRepository interface {
	Register(user User) (string, error)
	UpdateUserRole(email string, newRole string) error
	CheckUserExists(user User) (bool, User)
}

// Jwt service infrastructure interface
type IJwtService interface {
	GenerateToken(user_id primitive.ObjectID, userEmail string, userRole string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// password service infrastructure interface
type IPasswordService interface {
	HashPassword(userPassword string) (hashedPassword []byte, err error)
	CheckPassword(existingUserPassword string, userPassword string) error
}
