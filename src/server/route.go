package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wppzxc/tasks/src/server/handler"
	"github.com/wppzxc/tasks/src/server/static"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	registerRoute(e)
	return e
}

func registerRoute(e *echo.Echo) {
	// upload
	e.POST("/:username/uploads", static.UploadFiles)

	// serviceInfo
	e.GET("/:username/serviceInfo", static.GetServiceInfo)
	e.PUT("/:username/serviceInfo", static.UpdateServiceInfo)

	//users
	e.GET("/", handler.Test)
	e.GET("/users/:username", handler.GetUser)
	e.GET("/users/:username/:password", handler.CheckUserPassword)
	e.POST("/users/:username", handler.GenerateUser) // /userss/:username?parent=user_parent
	e.PUT("/users/:username", handler.UpdateUser)

	// tasks
	e.GET("/:username/tasks", handler.GetTasks)
	e.POST("/:username/tasks", handler.CreateTask)
	e.PUT("/:username/tasks", handler.UpdateTask)

	//commits
	e.GET("/:username/commits", handler.GetCommits)
	e.GET("/:username/:taskID/commits", handler.GetCurrentCommit)
	e.POST("/:username/:taskID/commits", handler.CreateCommits)
	// /:username/:taskID/commits?resolve=已通过 admin only
	e.PUT("/:username/:taskID/commits", handler.UpdateCommits)
}
