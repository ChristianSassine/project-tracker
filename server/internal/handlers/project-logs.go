package handlers

import (
	"net/http"
	"strconv"

	"github.com/krispier/projectManager/internal/services/db"

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

func GetProjectLogsByLimit(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Check if the limit query is specified
		logsLimit, ok := c.Request.URL.Query()["limit"]
		if ok {
			projectId, err := getProjectId(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

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
