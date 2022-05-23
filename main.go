package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Kaiya/todo/dao"
	"github.com/gin-gonic/gin"
)

var (
	sqlUser = flag.String("sqlUser", "root", "username for login mysql")
	sqlPass = flag.String("sqlPass", "", "password for mysql")
)

func main() {
	flag.Parse()
	err := SetFlagsFromEnvironment()
	if err != nil {
		panic(err)
	}
	if *sqlPass == "" {
		*sqlPass = os.Getenv("sqlPass")
	}

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

func SetFlagsFromEnvironment() (err error) {
	flag.VisitAll(func(f *flag.Flag) {
		name := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if value, ok := os.LookupEnv(name); ok {
			err2 := flag.Set(f.Name, value)
			if err2 != nil {
				err = fmt.Errorf("failed setting flag from environment: %w", err2)
			}
		}
	})

	return
}
