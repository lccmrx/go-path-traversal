package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func echoServer() {

	e := echo.New()

	e.GET("/hello/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	e.GET("/hello/world/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.HideBanner = true
	e.Start(":8001")
}
