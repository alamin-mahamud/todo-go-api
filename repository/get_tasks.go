package repository

import (
	"github.com/alaminmahamud/todo-app/configuration"
	"github.com/alaminmahamud/todo-app/usecase"
	"github.com/labstack/echo"
	"net/http"
)

var db = InitDB(configuration.DB_PATH)

func GetTasks() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			usecase.GetTasks(db),
			)
	}

}