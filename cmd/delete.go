package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/storage"
)

func DeleteTask(idStr string) {
	id, _ := strconv.Atoi(idStr)
	tasks, _ := storage.LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			storage.SaveTasks(tasks)
			fmt.Println("Task deleted 🗑️")
			return
		}
	}
	fmt.Println("Task not found")
}
