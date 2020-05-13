package handler

import (
	"net/http"

	"github.com/go_nuxt_2/server/api/api001/domain"
	"github.com/go_nuxt_2/server/environment/db"
	"github.com/labstack/echo"
)

// ShowTodos returns todos
func ShowTodos(c echo.Context) error {
	data, err := getTodos()

	if err != nil {
		return err
	}

	response := domain.ResponseData {
		Todos: data,
	}

	return c.JSON(http.StatusOK, response)

}

func getTodos() (tc []domain.Todo, err error) {
	db := db.CreateDBConnection() 
	defer db.Close()

	response := []domain.Todo{}

	err = db.Find(&response).Error 

	if err != nil {
		return response, err
	}

	return response, nil
}