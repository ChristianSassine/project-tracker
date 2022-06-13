package routes

import (
	"BugTracker/middlewares"
	"BugTracker/services/db"
	log "BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO : Might need refactoring
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

		urlQueries := c.Request.URL.Query()
		logsLimit, ok := urlQueries["limit"]
		// Get without any of the specified queries returns all the project logs
		if !ok {
			logs, err := db.GetAllLogs(projectId)
			if err != nil {
				log.PrintError(err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			c.JSON(http.StatusOK, logs)
			return
		}

		// Get the specified limit of recently added logs
		limit, err := strconv.Atoi(logsLimit[0])
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		tasks, err := db.GetLogsWithLimit(projectId, limit)
		c.JSON(http.StatusOK, tasks)
		return
	})
}
