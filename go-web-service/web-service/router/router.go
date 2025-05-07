package router

import (
	"task-management/web-service/controllers"
	"task-management/web-service/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.POST("register/", controllers.RegisterUser)
	router.POST("login/", controllers.LoginUser)

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("users/:id/", controllers.GetUserById)
		auth.GET("tasks/:id", controllers.GetTaskById)
		auth.POST("tasks/", controllers.AddTask)
		auth.GET("tasks/", controllers.GetAllTasks)
		auth.PUT("tasks/:id", controllers.UpdateTask)
		auth.DELETE("tasks/:id", controllers.DeleteTask)
	}

	return router
}
