package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func BasicHandler(c echo.Context) error {
	c.Logger().Info("hitting the context")
	return c.String(http.StatusOK, "Hello, World!")
}
