package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/wppzxc/tasks/src/database"
	"github.com/wppzxc/tasks/src/pkg/errors"
	"github.com/wppzxc/tasks/src/pkg/types"
	"github.com/wppzxc/tasks/src/pkg/utils"
	"github.com/wppzxc/tasks/src/server/service"
	"net/http"
	"strconv"
)

const (
	commitStatusStart  = "进行中"
	commitStatusCommit = "已提交"
	commitStatusDone   = "已通过"
	commitStatusRefuse = "被拒绝"
)

func GetCommits(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	state := ctx.QueryParam("state")
	var commits []*database.Commits
	var err error
	if username == types.AdminUser {
		commits, err = service.GetCommitsByName("", state)
	} else {
  		commits, err = service.GetCommitsByName(username, state)
	}
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("no commits")))
	}
	return ctx.JSON(http.StatusOK, commits)
}

func GetCurrentCommit(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	taskID := ctx.Param("taskID")
	if len(taskID) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("taskID must provide")))
	}
	commit, err := service.GetCurrentCommit(username, taskID)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, err))
	}
	return ctx.JSON(http.StatusOK, commit)
}

func CreateCommits(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	taskID := ctx.Param("taskID")
	if len(taskID) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("taskID must provide")))
	}
	task, err := service.GetTaskByID(taskID)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("task don't exist")))
	}
	task_id, _ := strconv.Atoi(taskID)
	commentKey := utils.GenerateKey(4)
	commit := &database.Commits{
		TaskID:     uint(task_id),
		TaskTitle:  task.Title,
		Username:   username,
		CommentKey: task.CommentBase + "(" + commentKey + ")",
		Status:     commitStatusStart,
	}
	newCommit, err := service.CreateCommit(commit)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	return ctx.JSON(http.StatusOK, newCommit)
}

func UpdateCommits(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}
	taskID := ctx.Param("taskID")
	if len(taskID) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("taskID must provide")))
	}
	commit := new(database.Commits)
	if err := ctx.Bind(commit); err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	resolve := ctx.QueryParam("resolve")
	if resolve == commitStatusDone {
		if username == types.AdminUser {
			if err := service.AddBonus(taskID, username); err != nil {
				fmt.Println(err)
			}
		} else {
			return ctx.JSON(errors.NewTaskServerError(http.StatusForbidden, fmt.Errorf("not admin user")))
		}
	}
	newCommit, err := service.UpdateCommits(commit)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	return ctx.JSON(http.StatusOK, newCommit)
}
