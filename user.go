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

// NewUser is create User Dao
func NewUser(id int, name string) User {
	return User{ID: id, Name: name}
}

// GetUser is return user response
func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user := User{ID: userID}
	has, err := engine.Get(&user)
	if has {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, err)
}

// GetUsers is Get user list
func GetUsers() []User {
	users := []User{}
	engine.Find(&users)
	return users
}
