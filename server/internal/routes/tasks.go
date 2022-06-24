package routes

import (
	"github.com/krispier/projectManager/internal/handlers"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func TasksRoutes(r *gin.RouterGroup, db *db.DB) {
	taskGroup := r.Group("", handlers.ValidateUserProject(db))

	// Getting tasks endpoint
	taskGroup.GET("/project/:projectId/tasks", handlers.GetTasksByLimit(db), handlers.GetTasksByState(db), handlers.GetAllTasks(db))

	// Getting tasks statistics endpoint
	taskGroup.GET("/project/:projectId/tasks/stats", handlers.GetTaskStats(db))

	// Creating a task endpoint
	taskGroup.POST("/project/:projectId/task", handlers.AddTask(db))

	// Updating a task endpoint
	taskGroup.PUT("/project/:projectId/task", handlers.UpdateTask(db))

	// Deleting a task endpoint
	taskGroup.DELETE("/project/:projectId/task", handlers.DeleteTask(db))

	// Updating a task position endpoint
	taskGroup.PATCH("/project/:projectId/task/position", handlers.UpdateTaskPosition(db))

	// Updating a task state endpoint
	taskGroup.PATCH("/project/:projectId/task/state", handlers.UpdateTaskState(db))
}
