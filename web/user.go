package web

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/gin-gonic/gin"
	"github.com/yuu-senoo/task-log/internal"
)

func CreateUser(context *gin.Context) {
	var user internal.User

	if err := context.BindJSON(&user); err != nil {
		slog.Error("message", err)
	}

	user = internal.CreateUser(user)

	context.IndentedJSON(http.StatusCreated, user)
}

func GetUsers(context *gin.Context) {
	users := internal.GetUsers()
	context.IndentedJSON(http.StatusOK, users)
}
