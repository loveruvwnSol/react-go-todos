package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []todo{
	{ID: 1, Title: "test todo 1"},
	{ID: 2, Title: "test todo 2"},
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/todos", getTodos)
	router.POST("/todos", postTodo)
	router.PUT("/todos:id", updateTodo)
	router.DELETE("/todos:id", deleteTodo)

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	var newTodo todo

	newTodo.ID = len(todos) + 1

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	var updateTodo todo

	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updateTodo.Title
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "success"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
}

// range
// gin.Context
