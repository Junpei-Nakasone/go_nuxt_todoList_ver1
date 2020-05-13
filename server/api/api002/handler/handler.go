package handler

import (
	"net/http"

	"github.com/go_nuxt_2/server/api/api002/domain"
	"github.com/go_nuxt_2/server/environment/db"
	"github.com/labstack/echo"
)

func AddTodo(c echo.Context) error {
	param := domain.Todo{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	if err := addTodo(param); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Created")

}

func addTodo(param domain.Todo) error {
	db := db.CreateDBConnection()

	data := domain.Todo{
		Name: param.Name,
	}
	err := db.Table("todos").Create(&data).Error

	return err
}
