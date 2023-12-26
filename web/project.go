package web

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/gin-gonic/gin"
	"github.com/yuu-senoo/task-log/internal"
)

func CreateProject(context *gin.Context) {
	var project internal.Project

	if err := context.BindJSON(&project); err != nil {
		slog.Error("message", err)
	}

	project = internal.CreateProject(project)

	context.IndentedJSON(http.StatusCreated, project)
}

func GetProjects(context *gin.Context) {
	projects := internal.GetProjects()
	context.IndentedJSON(http.StatusOK, projects)
}
