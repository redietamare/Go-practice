package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines routes.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize services and controllers
	taskService := data.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	// Define routes
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTaskByID)
		tasks.POST("", taskController.CreateTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}

	return router
}