package business_router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouting (e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Business!")
	})

	return e
}
