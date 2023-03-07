package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello World Function
func helloWord(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", helloWord)
	e.Logger.Fatal(e.Start(":8080"))
}
