package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//HelloWorld ...
func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
