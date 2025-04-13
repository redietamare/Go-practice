package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// TaskController handles task-related HTTP requests.
type TaskController struct {
	taskService *data.TaskService
}

// NewTaskController creates a new TaskController.
func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

// GetAllTasks returns all tasks.
func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks := c.taskService.GetAllTasks()
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTaskByID returns a task by its ID.
func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.taskService.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// CreateTask creates a new task.
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.taskService.CreateTask(newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// UpdateTask updates an existing task.
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.taskService.UpdateTask(id, updatedTask); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTask deletes a task by ID.
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.taskService.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}