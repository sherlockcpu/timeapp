package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Current date and time is: "+time.Now().String())
	})
	e.Logger.Fatal(e.Start(":80"))
}
