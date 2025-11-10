package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const TASK_DATA_DIR_ENV = "TASK_CLI_DATA_DIR"
const TASK_FILENAME = "tasks.json"
const TASK_DIR = ".task-tracker"

func getTasksFilePath() string {
	if dataDir := os.Getenv(TASK_DATA_DIR_ENV); dataDir != "" {
		if err := os.MkdirAll(dataDir, 0755); err == nil {
			return filepath.Join(dataDir, TASK_FILENAME)
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return TASK_FILENAME
	}

	taskDir := filepath.Join(homeDir, TASK_DIR)

	if err := os.MkdirAll(taskDir, 0755); err != nil {
		return TASK_FILENAME
	}

	return filepath.Join(taskDir, TASK_FILENAME)
}

var filename = getTasksFilePath()

func LoadTaskList() (*TaskList, error) {
	taskList := &TaskList{}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return taskList, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, taskList)
		if err != nil {
			return nil, fmt.Errorf("error parsing json: %v", err)
		}
	}

	return taskList, nil
}

func SaveTaskList(taskList *TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
