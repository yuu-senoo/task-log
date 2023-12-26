package internal

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ProjectID   uint   `json:"project_id" validate:"required"`
	Name        string `json:"name" validate:"required,max=256"`
	Description string `json:"description" validate:"max=512"`
}

func CreateTask(task Task) Task {
	if err := db.Create(&task).Error; err != nil {
		panic(err)
	}
	return task
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}
