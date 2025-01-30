package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	log.Print("TESTING APP")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
