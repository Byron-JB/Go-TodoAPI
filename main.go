package main

import (
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Priority    int    `json:"priority"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at,omitempty"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Go", Description: "Study the Go programming language", Priority: 1, Completed: false, CreatedAt: "2023-10-01T12:00:00Z"},
	{ID: 2, Title: "Build a REST API", Description: "Create a RESTful API using Go", Priority: 2, Completed: false, CreatedAt: "2023-10-02T12:00:00Z"},
	{ID: 3, Title: "Write Tests", Description: "Write unit tests for the API", Priority: 3, Completed: false, CreatedAt: "2023-10-03T12:00:00Z"},
	{ID: 4, Title: "Deploy Application", Description: "Deploy the Go application to a server", Priority: 4, Completed: false, CreatedAt: "2023-10-04T12:00:00Z"},
}

func GetListOfTodos(C *gin.Context) {
	C.JSON(200, todos)
}

func main() {

	fx.New().Run()

	//router := gin.Default()

	//router.GET("/todos", GetListOfTodos)

	//router.Run(":8080")
}
