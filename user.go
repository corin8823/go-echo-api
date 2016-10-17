package main

import (
	"fmt"

	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

// User is data type
type User struct {
	ID   int    `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/go_echo_api")
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// GetUser is return user response
func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	user := User{ID: userID}
	has, err := engine.Get(&user)
	if has {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, err)
}

// GetUsers is Get user list
func GetUsers(c echo.Context) error {
	users := []User{}
	engine.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// CreateUser is insert User
func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, "name is nil")
	}
	user := User{Name: name}
	_, err := engine.Insert(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, name)
}

// DeleteUser is User delete
func DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	user := User{ID: userID}
	_, deleteErr := engine.Delete(&user)
	if deleteErr != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.NoContent(http.StatusNoContent)
}
