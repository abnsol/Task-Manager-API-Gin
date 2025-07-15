package data

import (
	"errors"
	"task_management/models"
)

// In-memory storage for tasks
var tasks = []models.Task{
	{ID: "1", Title: "Go and Gin", Description: "Finish RestAPI basics with go and gin", Time: "12:30", Status: false},
	{ID: "2", Title: "PostMan", Description: "Start Project and Write Docs with post man", Time: "3:30", Status: false},
	{ID: "3", Title: "Mongo DB", Description: "Add persistency to the project", Time: "6:30", Status: false},
}

var errTask = errors.New("task not found")

// get all tasks
func GetTasks() []models.Task {
	return tasks
}

// returns error if task not found
func GetTaskById(id string) (models.Task, error) {

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errTask
}

// get task Index
func getIndex(newTask models.Task) (int, error) {
	for idx, task := range tasks {
		if newTask.ID == task.ID {
			return idx, nil
		}
	}

	return -1, errors.New("index not found")
}

func ReplaceTask(id string, newTask models.Task) (models.Task, error) {
	// not task with such id
	_, err := GetTaskById(id)
	if err != nil {
		return models.Task{}, errTask
	}

	idx, err := getIndex(newTask)
	if err != nil {
		return models.Task{}, errTask
	} else {
		tasks[idx] = newTask
		return newTask, nil
	}

}

func DeleteTask(id string) (string, error) {
	task, err := GetTaskById(id)
	if err != nil {
		return "", err
	} else {
		idx, err2 := getIndex(task)
		if err2 != nil {
			return "", errTask
		} else {
			tasks = append(tasks[:idx], tasks[idx+1:]...)
			return "Task Deleted", nil
		}
	}
}

// always creates Task so no need for error
func CreateTask(task models.Task) string {
	tasks = append(tasks, task)
	return "Task Created Successfully"
}
