package Controllers

import (
	"bookCRUD/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []Models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	task := Models.Task{AssingedTo: input.AssingedTo, Task: input.Task, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindTask(c *gin.Context) {
	var task Models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var task Models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput Models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book Models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
