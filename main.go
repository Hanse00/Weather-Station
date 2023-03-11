package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/:location", func(c echo.Context) error {
		location := c.Param("location")
		js, err := GetWeather(location)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, js)
	})
	e.Logger.Fatal(e.Start(":80"))
}
