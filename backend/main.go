package main

import (
	"app/app/handler"
	"app/app/model"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	var db = dbInit()
	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.User{})

	router.Use(middleware.CORSMiddleware())

	router.GET("/todos", handler.GetTodos(db))
	router.POST("/todos", handler.PostTodo(db))
	router.PUT("/todos:id", handler.UpdateTodo(db))
	router.DELETE("/todos:id", handler.DeleteTodo(db))

	router.GET("/user", middleware.AuthMiddleware(), handler.GetCurrentUser(db))

	router.POST("/signup", handler.SignUp(db))
	router.POST("/signin", handler.SignIn(db))

	router.Run(":8080")
}

func dbInit() *gorm.DB {
	dsn := "docker:docker@tcp(database)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
