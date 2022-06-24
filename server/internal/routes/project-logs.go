package routes

import (
	"github.com/krispier/projectManager/internal/handlers"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func ProjectLogsRoutes(r *gin.RouterGroup, db *db.DB) {
	logsGroup := r.Group("", handlers.ValidateUserProject(db))

	// Fetch the logs of the project
	logsGroup.GET("/project/:projectId/logs", handlers.GetProjectLogsByLimit(db), handlers.GetAllProjectLogs(db))
}