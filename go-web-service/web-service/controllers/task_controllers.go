package controllers

import (
	"net/http"
	"task-management/web-service/data"
	"task-management/web-service/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tasks})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := data.FindTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Task not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": task})
}

func AddTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "provide required fields!", "error": err.Error()})
		return
	}

	_, err := data.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Task created successfully"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "provide required fields!", "error": err.Error()})
		return
	}

	_, err := data.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	updatedTask, err := data.FindTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedTask})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Task deleted successfully"})
}
