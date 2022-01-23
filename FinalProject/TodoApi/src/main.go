package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"host.local/go/golang-todo-api/src/database"
	"host.local/go/golang-todo-api/src/handlers"
)

func main() {
	log.Info("Starting the application")
	database.Init()

	todoHandler := handlers.NewTodo()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.GET("/api/todos", todoHandler.GetTodos)

	router.POST("/api/todos", todoHandler.CreateTodo)

	router.PUT("/api/todos/:id", todoHandler.UpdateTodo)

	router.DELETE("/api/todos/:id", todoHandler.DeleteTodo)

	router.Run(":9090")
}
