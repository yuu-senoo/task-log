package internal

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateTask(task Task) Task {
	db.Create(&task)
	return task
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}
