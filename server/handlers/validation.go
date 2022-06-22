package handlers

import (
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt-token"
	"BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Checking if the token is valid on every entry
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if err := jwtToken.ValidateToken(token, false); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Next()
	}
}

// Validating that user has permissions to access the project's tasks
func ValidateUserProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenInfo, err := jwtToken.ExtractClaims(token)
		if err != nil {
			utilities.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(tokenInfo.Subject)
		if err != nil {
			utilities.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projectId, err := strconv.Atoi(c.Param("projectId"))
		if err != nil {
			utilities.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !db.IsUserInProject(userId, projectId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
