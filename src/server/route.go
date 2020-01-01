package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wppzxc/tasks/src/server/handler"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	registerRoute(e)
	return e
}

func registerRoute(e *echo.Echo) {
	e.GET("/", handler.Test)
	e.GET("/users/:username", handler.GetUser)
	e.GET("/users/:username/:password", handler.CheckUserPassword)
	// /userss/:username?parent=user_parent
	e.POST("/users/:username", handler.GenerateUser)
	e.PUT("/users/:username", handler.UpdateUser)
}
