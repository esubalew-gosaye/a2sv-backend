package router

import (
	"task-management/web-service/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.GET("tasks/", controllers.GetAllTasks)
	router.GET("tasks/:id", controllers.GetTaskById)

	router.POST("tasks/", controllers.AddTask)
	router.PUT("tasks/:id", controllers.UpdateTask)
	router.DELETE("tasks/:id", controllers.DeleteTask)

	return router
}
