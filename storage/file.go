package storage

import (
	"encoding/json"
	"os"
	"todo-cli/models"
)

const fileName = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
