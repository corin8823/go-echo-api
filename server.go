package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/users/:userID", GetUser)
	e.GET("/users", GetUsers)
	e.POST("/users", CreateUser)
	e.DELETE("/users/:userID", DeleteUser)
	e.Run(standard.New(":8080"))
}
