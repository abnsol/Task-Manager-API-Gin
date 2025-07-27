package domain

import (
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
type TaskUseCase interface {
	CreateTask(task Task) (string, error)
	GetTasks() []Task
	GetTaskById(id string) (Task, error)
	ReplaceTask(id string, newTask Task) (Task, error)
	DeleteTask(id string) (string, error)
}

// Task repository interfaces
type TaskRepository interface {
	CreateTask(task Task) (string, error)
	GetTasks() []Task
	GetTaskById(id string) (Task, error)
	ReplaceTask(id string, newTask Task) (Task, error)
	DeleteTask(id string) (string, error)
}

// User use case interfaces
type UserUseCase interface {
	Register(user User) (string, error)
	Login(user User) (string, string, error)
	UpdateUserRole(email string, newRole string) error
}

// User repository interfaces
type UserRepository interface {
	Register(user User) (string, error)
	Login(user User) (string, string, error)
	UpdateUserRole(email string, newRole string) error
}
