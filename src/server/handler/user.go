package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/wppzxc/tasks/src/database"
	"github.com/wppzxc/tasks/src/server/service"
	"net/http"
)

func Test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "test")
}

func GetUser(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(http.StatusNotFound, "username must provide")
	}
	user, err := service.GetUserByName(username)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func CheckUserPassword(ctx echo.Context) error {
	username := ctx.Param("username")
	password := ctx.Param("password")
	user, err := service.GetUserByNameAndPassword(username, password)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func GenerateUser(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(http.StatusNotFound, "username must provide")
	}
	user := &database.Users{
		Name:     username,
		Password: username,
		Status:   database.UserNormal,
	}
	parent := ctx.QueryParam("parent")
	if len(parent) > 0 {
		_, err := service.GetUserByName(parent)
		if err == nil {
			user.Parent = parent
		}
	}
	user, err := service.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx echo.Context) error {

	register := ctx.QueryParam("register")
	if register == "true" || register == "True" {
		return RegisterUser(ctx)
	}

	user := new(database.Users)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	newUser, err := service.UpdateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, newUser)
}

func RegisterUser(ctx echo.Context) error {
	user := new(database.Users)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if len(user.AlipayAccount) == 0 {
		return ctx.JSON(http.StatusForbidden, fmt.Errorf("must provide alipayAccount"))
	}
	user.Registered = true
	oldUser, err := service.GetUserByName(user.Name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, fmt.Errorf("Error in get user : %s, '%s' ", user.Name, err))
	}
	if oldUser.Registered {
		return ctx.JSON(http.StatusConflict, fmt.Errorf("Error in register user : %s, '%s' ", user.Name, "user has been registered"))
	}
	newUser, err := service.UpdateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if len(newUser.Parent) > 0 {
		if err := service.UpdateParentsByRegister(newUser.Parent); err != nil {
			ctx.Logger().Error(err)
		}
	}
	return ctx.JSON(http.StatusOK, newUser)
}
