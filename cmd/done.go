package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/storage"
)

func CompleteTask(idStr string) {
	id, _ := strconv.Atoi(idStr)
	tasks, _ := storage.LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			storage.SaveTasks(tasks)
			fmt.Println("Task is done ✅")
			return
		}
	}
	fmt.Println("Task not found.")
}
