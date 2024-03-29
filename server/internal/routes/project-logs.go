package routes

import (
	"github.com/ChristianSassine/projectManager/internal/handlers"
	"github.com/ChristianSassine/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func ProjectLogsRoutes(r *gin.RouterGroup, db *db.DB) {
	logsGroup := r.Group("/project/:projectId", handlers.ValidateUserProject(db))

	// Fetch the logs of the project
	logsGroup.GET("/logs", handlers.GetProjectLogsByLimit(db), handlers.GetAllProjectLogs(db))
}
