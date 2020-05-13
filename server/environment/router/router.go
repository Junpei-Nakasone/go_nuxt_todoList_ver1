package router

import (
	"net/http"

	api001 "github.com/go_nuxt_2/server/api/api001/handler"
	api002 "github.com/go_nuxt_2/server/api/api002/handler"
	api003 "github.com/go_nuxt_2/server/api/api003/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter returns router
func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/todos", api001.ShowTodos)
	e.POST("/newTodo", api002.AddTodo)
	e.POST("/editTodo", api003.EditTodo)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "THIS IS TEST.")
	})
	e.GET("/hello2", func(c echo.Context) error {
		return c.String(http.StatusOK, "THIS IS HELLO2")
	})
	e.Use(middleware.CORS())

	return e
}
