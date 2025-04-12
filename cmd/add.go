package cmd

import (
	"fmt"
	"todo-cli/models"
	"todo-cli/storage"
)

func AddTask(title string) { 
	tasks, _ := storage.LoadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task := models.Task{ID: id, Title: title}
	tasks = append(tasks, task)
	storage.SaveTasks(tasks)
	fmt.Println("Task added:", title)
}
