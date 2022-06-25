package routes

import (
	"github.com/ChristianSassine/projectManager/internal/handlers"
	"github.com/ChristianSassine/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func TaskCommentsRoutes(r *gin.RouterGroup, db *db.DB) {
	commentsGroup := r.Group("/project/:projectId/task/:taskId", handlers.ValidateUserProject(db))

	// Get a task's comments endpoint
	commentsGroup.GET("/comments", handlers.GetAllComments(db))

	// Add a comment to a task endpoint
	commentsGroup.POST("/comment", handlers.AddComment(db))
}
