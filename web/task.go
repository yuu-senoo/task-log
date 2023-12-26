package web

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/gin-gonic/gin"
	"github.com/yuu-senoo/task-log/internal"
)

func CreateTask(context *gin.Context) {
	var task internal.Task

	if err := context.BindJSON(&task); err != nil {
		slog.Error("message", err)
	}

	task = internal.CreateTask(task)

	context.IndentedJSON(http.StatusCreated, task)
}

func GetTasks(context *gin.Context) {
	tasks := internal.GetTasks()
	context.IndentedJSON(http.StatusOK, tasks)
}
