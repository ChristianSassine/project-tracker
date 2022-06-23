package routes

import (
	"BugTracker/handlers"
	"BugTracker/services/db"

	"github.com/gin-gonic/gin"
)

func TaskCommentsRoutes(r *gin.RouterGroup, db *db.DB) {
	commentsGroup := r.Group("", handlers.ValidateUserProject(db))

	// Get a task's comments endpoint
	commentsGroup.GET("/project/:projectId/task/:taskId/comments", handlers.GetAllComments(db))

	// Add a comment to a task endpoint
	commentsGroup.POST("/project/:projectId/task/:taskId/comment", handlers.AddComment(db))
}
