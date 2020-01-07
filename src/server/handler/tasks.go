package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/wppzxc/tasks/src/database"
	"github.com/wppzxc/tasks/src/pkg/errors"
	"github.com/wppzxc/tasks/src/pkg/types"
	"github.com/wppzxc/tasks/src/server/service"
	"net/http"
)

func GetTasks(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	task, err := service.GetCurrentTasks()
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, err))
	}
	return ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 || username != types.AdminUser {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	task := &database.Tasks{}
	if err := ctx.Bind(task); err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	newTask, err := service.CreateTask(task)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	return ctx.JSON(http.StatusOK, newTask)
}

func UpdateTask(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 || username != types.AdminUser {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	task := &database.Tasks{}
	if err := ctx.Bind(task); err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	if task.ID <= 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("invaild taskID")))
	}
	newTask, err := service.UpdateTask(task)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	return ctx.JSON(http.StatusOK, newTask)
}
