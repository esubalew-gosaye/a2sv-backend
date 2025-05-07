package controllers

import (
	"net/http"
	"task-management/web-service/data"
	"task-management/web-service/models"
	"task-management/web-service/utils"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := data.LoginUser(loginData)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful", "token": token})
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := data.AddUser(user)
	if err != nil {
		c.JSON(500, gin.H{"message": "We can't add new user"})
	}

	c.JSON(200, gin.H{"message": "User registered successfully", "data": resp})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := data.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "User not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}
