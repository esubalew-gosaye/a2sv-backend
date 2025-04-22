package data

import (
	"task-management/web-service/models"
	"time"
)

var tasks = []models.Task{
	{Id: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{Id: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{Id: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task {
	return tasks
}

func AddTask(task models.Task) []models.Task {
	tasks = append(tasks, task)
	return tasks
}

func FindTaskById(id string) (int, *models.Task) {
	for i, task := range tasks {
		if task.Id == id {
			return i, &task
		}
	}
	return -1, nil
}

func UpdateTask(task models.Task) *models.Task {
	for i, t := range tasks {
		if task.Id == t.Id {
			tasks[i] = task
			return &tasks[i]
		}
	}
	return nil
}

func DeleteTask(id string) {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}
