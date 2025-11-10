package main

import "time"

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

func (ts TaskStatus) IsValid() bool {
	switch ts {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return true
	}

	return false
}

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(id int, desc string) Task {
	now := time.Now()

	return Task{
		ID:          id,
		Description: desc,
		Status:      TaskStatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (tl *TaskList) GetNextID() int {
	maxId := 0

	for _, task := range tl.Tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	return maxId + 1
}

func (tl *TaskList) AddTask(desc string) Task {
	newTask := NewTask(tl.GetNextID(), desc)

	tl.Tasks = append(tl.Tasks, newTask)

	return newTask
}

func (tl *TaskList) UpdateTask(id int, desc string) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Description = desc
			tl.Tasks[i].UpdatedAt = time.Now()
			return true
		}
	}

	return false
}

func (tl *TaskList) DeleteTask(id int) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return true
		}
	}

	return false
}

func (tl *TaskList) UpdateTaskStatus(id int, status TaskStatus) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Status = status
			tl.Tasks[i].UpdatedAt = time.Now()
			return true
		}
	}

	return false
}

func (tl *TaskList) GetTasksByStatus(status TaskStatus) []Task {
	var tasks []Task

	for _, task := range tl.Tasks {
		if task.Status == status || status == "" {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
