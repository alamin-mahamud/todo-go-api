package app

import (
	"github.com/alaminmahamud/todo-app/repository"
	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()
	e.Static("/", "public")
	e.GET("/tasks", repository.GetTasks())
	e.Start(":8080")
}