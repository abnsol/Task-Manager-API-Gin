package usecases

import domain "task_management/Domain"

type TaskUseCase struct {
	TaskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepository domain.TaskRepository) domain.TaskUseCase {
	return &TaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (tu *TaskUseCase) CreateTask(task domain.Task) (string, error) {
	return tu.TaskRepository.CreateTask(task)
}

func (tu *TaskUseCase) GetTasks() []domain.Task {
	return tu.TaskRepository.GetTasks()
}

func (tu *TaskUseCase) GetTaskById(id string) (domain.Task, error) {
	return tu.TaskRepository.GetTaskById(id)
}

func (tu *TaskUseCase) ReplaceTask(id string, newTask domain.Task) (domain.Task, error) {
	return tu.TaskRepository.ReplaceTask(id, newTask)
}

func (tu *TaskUseCase) DeleteTask(id string) (string, error) {
	return tu.TaskRepository.DeleteTask(id)
}
