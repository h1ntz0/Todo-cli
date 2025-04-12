package main

import (
	"fmt"
	"os"
	"todo-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You can use the following commands: add <task>, list, done <id>, delete <id>")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("You don't fill the task")
			return
		}
		cmd.AddTask(os.Args[2])
	case "list":
		cmd.ListTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Which ID sir?")
			return
		}
		cmd.CompleteTask(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Which ID sir?")
			return
		}
		cmd.DeleteTask(os.Args[2])
	default:
		fmt.Println("wrong command")
	}
}
