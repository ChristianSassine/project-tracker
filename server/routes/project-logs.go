package routes

import (
	"BugTracker/handlers"
	"BugTracker/services/db"

	"github.com/gin-gonic/gin"
)

func ProjectLogsRoutes(r *gin.RouterGroup, db *db.DB) {
	logsGroup := r.Group("", handlers.ValidateUserProject(db))

	// Fetch the logs of the project
	logsGroup.GET("/project/:projectId/logs", handlers.GetLimitedProjectLogs(db), handlers.GetAllProjectLogs(db))
}
