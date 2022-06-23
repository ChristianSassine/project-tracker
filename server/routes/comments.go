package routes

import (
	"BugTracker/handlers"
	"BugTracker/services/db"

	"github.com/gin-gonic/gin"
)

func TaskCommentsRoutes(r *gin.RouterGroup, db *db.DB) {
	taskGroup := r.Group("", handlers.ValidateUserProject(db))

	// Get a task's comments endpoint
	taskGroup.GET("/project/:projectId/task/:taskId/comments", handlers.GetAllComments(db))

	// Add a comment to a task endpoint
	taskGroup.POST("/project/:projectId/task/:taskId/comment", handlers.AddComment(db))
}
