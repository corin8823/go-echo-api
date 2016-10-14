package main

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:userID", func(c echo.Context) error {
		userID, _ := strconv.Atoi(c.Param("userID"))
		user := FindUserByID(userID)
		return c.String(http.StatusOK, "Hello World "+user.Name)
	})
	e.Run(standard.New(":8080"))
}
