package main

import (
	"todo-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/todos", handlers.GetTodos)
		v1.POST("todos", handlers.CreateTodo)
		v1.PUT("/todos/:id", handlers.UpdateTodo)
		v1.DELETE("/todos/:id", handlers.DeleteTodo)

	}
	r.Run(":8080")
}
