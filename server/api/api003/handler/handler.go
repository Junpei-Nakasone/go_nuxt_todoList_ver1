package handler

import (
	"net/http"

	"github.com/go_nuxt_2/server/api/api003/domain"
	"github.com/go_nuxt_2/server/environment/db"
	"github.com/labstack/echo"
)

func EditTodo(c echo.Context) error {
	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	if err := editTodo(param); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "Updated")

}

func editTodo(param domain.RequestParam) error {
	db := db.CreateDBConnection()

	updateData := domain.UpdateData{
		Name: param.NewTodo,
	}
	err := db.Table("todos").
		Where("id = ?", param.ID).
		UpdateColumns(updateData).Error

	if err != nil {
		return err
	}

	return nil
}
