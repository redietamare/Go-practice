package data

import (
	"errors"
	"task_manager/models"
	"time"
	// "github.com/gin-gonic/gin"	
)

// TaskService manages the in-memory task data.
type TaskService struct {
	tasks []models.Task
}

// NewTaskService initializes a new TaskService with mock data.
func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{
			{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
			{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
			{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
		},
	}
}

// GetAllTasks returns all tasks.
func (s *TaskService) GetAllTasks() []models.Task {
	return s.tasks
}

// GetTaskByID returns a task by its ID or an error if not found.
func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// CreateTask adds a new task.
func (s *TaskService) CreateTask(task models.Task) {
	s.tasks = append(s.tasks, task)
}

// UpdateTask updates a task by ID with the provided task data.
func (s *TaskService) UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range s.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				s.tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				s.tasks[i].Description = updatedTask.Description
			}
			return nil
		}
	}
	return errors.New("task not found")
}

// DeleteTask removes a task by ID.
func (s *TaskService) DeleteTask(id string) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}