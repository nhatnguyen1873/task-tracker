package main

import "fmt"

func printUsage() {
	usage := `
Task Tracker - Manage your tasks efficiently

Usage:
  task-cli <command> [arguments]

Commands:
  add <description>              Add a new task
  update <id> <description>      Update an existing task
  delete <id>                    Delete a task
  mark-in-progress <id>          Mark a task as in progress
  mark-done <id>                 Mark a task as done
  list                           List all tasks
  list todo                      List all tasks that are not done
  list in-progress               List all tasks that are in progress
  list done                      List all tasks that are done

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy groceries and cook dinner"
  task-cli mark-in-progress 1
  task-cli mark-done 1
  task-cli list
  task-cli list done
`
	fmt.Println(usage)
}

func printHelp() {
	fmt.Println("Run 'task-cli help' for usage instructions")
}
