package usecases

import (
	"errors"
	"testing"

	domain "task_management/Domain"
	mock_domain "task_management/Mocks/domain"

	"github.com/stretchr/testify/suite"
)

// TaskUseCaseSuite defines the test suite for TaskUseCase
type TaskUseCaseSuite struct {
	suite.Suite
	mockTaskRepo *mock_domain.ITaskRepository
	taskUseCase  domain.ITaskUseCase
}

// SetupTest is run before each test in the suite
func (s *TaskUseCaseSuite) SetupTest() {
	s.mockTaskRepo = new(mock_domain.ITaskRepository)
	s.taskUseCase = NewTaskUseCase(s.mockTaskRepo)
}

// TestTaskUseCaseSuite runs the entire test suite
func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}

// TestCreateTaskSuccess tests successful task creation
func (s *TaskUseCaseSuite) TestCreateTaskSuccess() {
	task := domain.Task{Title: "Test Task"}
	s.mockTaskRepo.On("CreateTask", task).Return(nil).Once()

	msg, err := s.taskUseCase.CreateTask(task)

	s.NoError(err)
	s.Equal("Task Successfully created", msg)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestCreateTaskFailure tests task creation failure
func (s *TaskUseCaseSuite) TestCreateTaskFailure() {
	task := domain.Task{Title: "Test Task"}
	s.mockTaskRepo.On("CreateTask", task).Return(errors.New("db error")).Once()

	_, err := s.taskUseCase.CreateTask(task)

	s.Error(err)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestGetTasks tests fetching all tasks
func (s *TaskUseCaseSuite) TestGetTasks() {
	expectedTasks := []domain.Task{
		{ID: "1", Title: "Task 1"},
		{ID: "2", Title: "Task 2"},
	}
	s.mockTaskRepo.On("GetTasks").Return(expectedTasks).Once()

	tasks := s.taskUseCase.GetTasks()

	s.Equal(expectedTasks, tasks)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestGetTaskByIdSuccess tests fetching a task by ID successfully
func (s *TaskUseCaseSuite) TestGetTaskByIdSuccess() {
	taskID := "task-123"
	expectedTask := domain.Task{ID: taskID, Title: "My Task"}
	s.mockTaskRepo.On("GetTaskById", taskID).Return(expectedTask, nil).Once()

	task, err := s.taskUseCase.GetTaskById(taskID)

	s.NoError(err)
	s.Equal(expectedTask, task)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestGetTaskByIdNotFound tests fetching a non-existent task
func (s *TaskUseCaseSuite) TestGetTaskByIdNotFound() {
	taskID := "task-123"
	s.mockTaskRepo.On("GetTaskById", taskID).Return(domain.Task{}, errors.New("not found")).Once()

	_, err := s.taskUseCase.GetTaskById(taskID)

	s.Error(err)
	s.Equal("task not found", err.Error())
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestReplaceTaskSuccess tests replacing a task successfully
func (s *TaskUseCaseSuite) TestReplaceTaskSuccess() {
	taskID := "task-123"
	newTask := domain.Task{Title: "Updated Task"}
	expectedTask := domain.Task{ID: taskID, Title: "Updated Task"}
	s.mockTaskRepo.On("ReplaceTask", taskID, newTask).Return(expectedTask, nil).Once()

	task, err := s.taskUseCase.ReplaceTask(taskID, newTask)

	s.NoError(err)
	s.Equal(expectedTask, task)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestReplaceTaskFailure tests failing to replace a task
func (s *TaskUseCaseSuite) TestReplaceTaskFailure() {
	taskID := "task-123"
	newTask := domain.Task{Title: "Updated Task"}
	s.mockTaskRepo.On("ReplaceTask", taskID, newTask).Return(domain.Task{}, errors.New("db error")).Once()

	_, err := s.taskUseCase.ReplaceTask(taskID, newTask)

	s.Error(err)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestDeleteTaskSuccess tests deleting a task successfully
func (s *TaskUseCaseSuite) TestDeleteTaskSuccess() {
	taskID := "task-123"
	s.mockTaskRepo.On("DeleteTask", taskID).Return(nil).Once()

	msg, err := s.taskUseCase.DeleteTask(taskID)

	s.NoError(err)
	s.Equal("Task deleted successfully", msg)
	s.mockTaskRepo.AssertExpectations(s.T())
}

// TestDeleteTaskNotFound tests deleting a non-existent task
func (s *TaskUseCaseSuite) TestDeleteTaskNotFound() {
	taskID := "task-123"
	s.mockTaskRepo.On("DeleteTask", taskID).Return(errors.New("not found")).Once()

	_, err := s.taskUseCase.DeleteTask(taskID)

	s.Error(err)
	s.Equal("task not found", err.Error())
	s.mockTaskRepo.AssertExpectations(s.T())
}
