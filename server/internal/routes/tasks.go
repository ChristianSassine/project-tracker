package routes

import (
	"github.com/ChristianSassine/projectManager/internal/handlers"
	"github.com/ChristianSassine/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func TasksRoutes(r *gin.RouterGroup, db *db.DB) {
	taskGroup := r.Group("/project/:projectId", handlers.ValidateUserProject(db))

	// Getting tasks endpoint
	taskGroup.GET("/tasks", handlers.GetTasksByLimit(db), handlers.GetTasksByState(db), handlers.GetAllTasks(db))

	// Getting tasks statistics endpoint
	taskGroup.GET("/tasks/stats", handlers.GetTaskStats(db))

	// Creating a task endpoint
	taskGroup.POST("/task", handlers.AddTask(db))

	// Updating a task endpoint
	taskGroup.PUT("/task", handlers.UpdateTask(db))

	// Deleting a task endpoint
	taskGroup.DELETE("/task", handlers.DeleteTask(db))

	// Updating a task position endpoint
	taskGroup.PATCH("/task/position", handlers.UpdateTaskPosition(db))

	// Updating a task state endpoint
	taskGroup.PATCH("/task/state", handlers.UpdateTaskState(db))
}
