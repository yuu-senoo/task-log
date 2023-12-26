package internal

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks" gorm:"foreignKey:ProjectID" validate:"-"`
}

func CreateProject(project Project) Project {
	db.Create(&project)
	return project
}

func GetProjects() []Project {
	var projects []Project
	db.Find(&projects)
	return projects
}
