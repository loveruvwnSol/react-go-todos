package handler

import (
	"app/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos = []model.Todo{}
		result := db.Find(&todos)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not read todos"})
			return
		}

		c.IndentedJSON(http.StatusOK, &todos)
	}
}

func FetchTodos(db *gorm.DB) ([]model.Todo, error) {
	var todos = []model.Todo{}
	result := db.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func PostTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo model.Todo

		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		result := db.Create(&newTodo)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not create todo"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Can not create todo"})
			return
		}

		todos, err := FetchTodos(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not read todos"})
			return
		}
		c.IndentedJSON(http.StatusCreated, &todos)
	}
}

func UpdateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		var updateTodo model.Todo

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := c.ShouldBind(&updateTodo); err != nil {
			return
		}

		result := db.Model(&model.Todo{}).Where("id = ?", id).Update("title", updateTodo.Title)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not update todo"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found the todo"})
			return
		}

		todos, err := FetchTodos(db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not read todos"})
			return
		}
		c.IndentedJSON(http.StatusOK, todos)
	}
}

func DeleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		result := db.Where("id = ?", id).Delete(&model.Todo{})

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not delete todo"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found the todo"})
			return
		}

		todos, err := FetchTodos(db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Can not read todos"})
			return
		}
		c.IndentedJSON(http.StatusOK, &todos)
	}
}
