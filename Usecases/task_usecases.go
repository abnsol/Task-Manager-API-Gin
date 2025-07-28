package usecases

import (
	"errors"
	domain "task_management/Domain"
)

type TaskUseCase struct {
	TaskRepository domain.ITaskRepository
}

func NewTaskUseCase(taskRepository domain.ITaskRepository) domain.ITaskUseCase {
	return &TaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (tu *TaskUseCase) CreateTask(task domain.Task) (string, error) {
	err := tu.TaskRepository.CreateTask(task)

	if err != nil {
		return "", err
	}
	return "Task Successfully created", nil
}

func (tu *TaskUseCase) GetTasks() []domain.Task {
	return tu.TaskRepository.GetTasks()

}

func (tu *TaskUseCase) GetTaskById(id string) (domain.Task, error) {
	task, err := tu.TaskRepository.GetTaskById(id)

	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (tu *TaskUseCase) ReplaceTask(id string, newTask domain.Task) (domain.Task, error) {
	return tu.TaskRepository.ReplaceTask(id, newTask)
}

func (tu *TaskUseCase) DeleteTask(id string) (string, error) {
	err := tu.TaskRepository.DeleteTask(id)
	if err != nil {
		return "", errors.New("task not found")
	}
	return "Task deleted successfully", nil
}
