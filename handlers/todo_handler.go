package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"todo-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var todos = []models.Todo{}
var nextID int64

func nextTodoID() string {
	n := atomic.AddInt64(&nextID, 1)
	return fmt.Sprintf("%d", n)
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = uuid.NewString()
	todos = append(todos, input)

	c.JSON(http.StatusCreated, input)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var input models.Todo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = input.Title
			todos[i].Desc = input.Desc
			todos[i].Completed = input.Completed
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}
