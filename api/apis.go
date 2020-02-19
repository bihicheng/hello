package api

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
    webcontent = "Hello Everyone!"
)

// Hello ...
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(webcontent))
}
