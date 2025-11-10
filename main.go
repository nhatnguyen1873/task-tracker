package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	taskList, err := LoadTaskList()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(taskList, os.Args[2:])
	case "list":
		handleList(taskList, os.Args[2:])
	case "update":
		handleUpdate(taskList, os.Args[2:])
	case "delete":
		handleDelete(taskList, os.Args[2:])
	case "mark-done":
		handleMarkDone(taskList, os.Args[2:])
	case "mark-in-progress":
		handleMarkInProgress(taskList, os.Args[2:])
	case "help":
		printUsage()
	default:
		fmt.Printf("Invalid command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func handleAdd(taskList *TaskList, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: add <description>")
		os.Exit(1)
	}

	task := taskList.AddTask(args[0])

	err := SaveTaskList(taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func handleList(taskList *TaskList, args []string) {
	var status string

	if len(args) > 0 {
		status = args[0]
	}

	if status != "" && !TaskStatus(status).IsValid() {
		fmt.Println("Error: invalid status (todo, in-progress, done)")
		os.Exit(1)
	}

	actualStatus := TaskStatus(status)
	tasks := taskList.GetTasksByStatus(actualStatus)

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("ID | Description | Status | Created At | Updated At")
	fmt.Println("-- | ----------- | ------ | ---------- | ----------")

	for _, task := range tasks {
		fmt.Printf("%d | %s | %s | %s | %s\n",
			task.ID,
			task.Description,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

func handleUpdate(taskList *TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: update <id> <description>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error parsing ID: %v\n", err)
		os.Exit(1)
	}

	if !taskList.UpdateTask(id, args[1]) {
		fmt.Printf("Task not found (ID: %d)\n", id)
		os.Exit(1)
	}

	err = SaveTaskList(taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task updated successfully (ID: %d)\n", id)
}

func handleDelete(taskList *TaskList, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: delete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error parsing ID: %v\n", err)
		os.Exit(1)
	}

	if !taskList.DeleteTask(id) {
		fmt.Printf("Task not found (ID: %d)\n", id)
		os.Exit(1)
	}

	err = SaveTaskList(taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task deleted successfully (ID: %v)\n", id)
}

func handleMarkDone(taskList *TaskList, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mark-done <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error parsing ID: %v\n", err)
		os.Exit(1)
	}

	if !taskList.UpdateTaskStatus(id, TaskStatusDone) {
		fmt.Printf("Task not found (ID: %d)\n", id)
		os.Exit(1)
	}

	err = SaveTaskList(taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task marked as done (ID: %d)\n", id)
}

func handleMarkInProgress(taskList *TaskList, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: mark-in-progress <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error parsing ID: %v\n", err)
		os.Exit(1)
	}

	if !taskList.UpdateTaskStatus(id, TaskStatusInProgress) {
		fmt.Printf("Task not found (ID: %d)\n", id)
		os.Exit(1)
	}

	err = SaveTaskList(taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task marked as in-progress (ID: %d)\n", id)
}
