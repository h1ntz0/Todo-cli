package cmd

import (
	"fmt"
	"todo-cli/storage"
)

func ListTasks() {
	tasks, _ := storage.LoadTasks()
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found.")
		return
	}
	for _, t := range tasks {
		status := "❌"
		if t.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
	}
}
