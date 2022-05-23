package main

import (
	"flag"

	"github.com/Kaiya/todo/dao"
	"github.com/gin-gonic/gin"
)

var (
	sqlUser = flag.String("sqlUser", "root", "username for login mysql")
	sqlPass = flag.String("sqlPass", "", "password for mysql")
)

func main() {
	flag.Parse()
	if *sqlUser == "" || *sqlPass == "" {
		panic("sqlPass is empty")
	}
	dao.Init(*sqlUser, *sqlPass)
	router := gin.Default()
	api := router.Group("/api/v1/todos")
	{
		api.POST("/", dao.CreateTodo)
		api.GET("/", dao.FetchAllTodo)
		api.GET("/:id", dao.FetchSingleTodo)
		api.PUT("/:id", dao.UpdateTodo)
		api.DELETE("/:id", dao.DeleteTodo)
	}
	router.Run()
}
