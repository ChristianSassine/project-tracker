package handlers

import (
	"BugTracker/services/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProjectLogs(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		logs, err := db.GetAllLogs(projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, logs)
		return

	}
}

func GetLimitedProjectLogs(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Check if the limit query is specified
		urlQueries := c.Request.URL.Query()
		logsLimit, ok := urlQueries["limit"]
		if ok {
			// Get the specified limit of recently added logs
			limit, err := strconv.Atoi(logsLimit[0])
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			logs, err := db.GetLogsWithLimit(projectId, limit)
			c.AbortWithStatusJSON(http.StatusOK, logs)
		}
	}
}
