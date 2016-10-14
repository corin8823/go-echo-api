package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
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

// FindUserByID is Get User
func FindUserByID(id int) User {
	user := User{}
	engine.ID(id).Get(&user)
	return user
}
