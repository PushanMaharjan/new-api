package Routes

import (
	"bookCRUD/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/tasks", Controllers.FindTasks)
	r.POST("/tasks", Controllers.CreateTask)
	r.GET("/tasks/:id", Controllers.FindTask)
	r.PATCH("/tasks/:id", Controllers.UpdateTask)
	r.DELETE("tasks/:id", Controllers.DeleteTask)
	return r
}
