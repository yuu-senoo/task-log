package internal

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(user User) User {
	db.Create(&user)
	return user
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
