package controllers

import (
	"net/http"
	"task-management/web-service/data"
	"task-management/web-service/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data.GetAllTasks()})
}

func GetTaskById(c *gin.Context) {
	id := c.Params.ByName("id")

	indx, task := data.FindTaskById(id)

	if indx != -1 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": task})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Task not Found!"})
}

func AddTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "provide required fields!", "error": err.Error()})
		return
	}

	data.AddTask(task)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data.GetAllTasks()})
}

func UpdateTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task

	indx, _ := data.FindTaskById(id)

	if indx == -1 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "provide required fields!", "error": err.Error()})
		return
	}

	task.Id = id
	data.UpdateTask(task)
	_, t := data.FindTaskById(id)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": t})
}

func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	indx, _ := data.FindTaskById(id)

	if indx == -1 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "task not found"})
		return
	}

	data.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data.GetAllTasks()})
}
