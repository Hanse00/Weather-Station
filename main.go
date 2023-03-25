package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	redis_addr := os.Getenv("REDIS_ADDR")
	if redis_addr == "" {
		redis_addr = "localhost:6379"
	}
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: redis_addr,
		Password: "",
		DB: 0,
	})
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/observation/:location", func(c echo.Context) error {
		location := c.Param("location")
		js, err := rdb.Get(ctx, location).Result()
		if err == redis.Nil {
			js, err = GetWeather(location)
			if err != nil {
				return err
			}
			err := rdb.Set(ctx, location, js, time.Minute).Err()
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, js)
		} else if err != nil {
			return err
		}
		return c.String(http.StatusOK, js)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
