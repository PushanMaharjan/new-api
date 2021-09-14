package main

import (
	"bookCRUD/Config"
	"bookCRUD/Controllers"
	"bookCRUD/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := Config.SetupDB()
	db.AutoMigrate(&Models.Task{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "milyoooooooo"})
	})
	r.GET("/tasks", Controllers.FindTasks)
	r.POST("/tasks", Controllers.CreateTask)
	r.GET("/tasks/:id", Controllers.FindTask)
	r.PATCH("/tasks/:id", Controllers.UpdateTask)
	r.DELETE("tasks/:id", Controllers.DeleteTask)
	r.Run()
}
