package main

import (
	"fmt"

	"golang.org/x/exp/slog"

	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yuu-senoo/task-log/config"
	"github.com/yuu-senoo/task-log/internal"
	"github.com/yuu-senoo/task-log/web"
)

func main() {
	// load config
	if err := godotenv.Load(); err != nil {
		slog.Info(".env file not found")
	}
	var cnf config.Config
	if err := env.Parse(&cnf); err != nil {
		panic(err)
	}
	slog.Debug(fmt.Sprintf("loaded config: %+v", cnf))

	// setup database
	internal.ConnectDB(cnf.DB)
	internal.MigrateDB()

	// setup router
	router := gin.Default()
	// user
	router.GET("/users", web.GetUsers)
	router.POST("/user", web.CreateUser)
	// project
	router.GET("/projects", web.GetProjects)
	router.POST("/project", web.CreateProject)
	// task
	router.GET("/tasks", web.GetTasks)
	router.POST("/task", web.CreateTask)

	router.Run(fmt.Sprintf("localhost:%d", cnf.Port))
}
