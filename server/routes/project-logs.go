package routes

import (
	"BugTracker/middlewares"
	"BugTracker/services/db"
	log "BugTracker/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectLogsRoutes(r *gin.RouterGroup, db *db.DB) {
	logsGroup := r.Group("", middlewares.ValidUserProjectAccessMiddleware(db))

	// Fetch the logs of the project
	logsGroup.GET("/project/:projectId/logs", func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		logs, err := db.GetAllLogs(projectId)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, logs)
	})
}
