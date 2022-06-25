package handlers

import (
	"net/http"
	"strconv"

	projectErrors "github.com/ChristianSassine/projectManager/internal/common/project-errors"
	"github.com/ChristianSassine/projectManager/internal/services/db"
	jwtToken "github.com/ChristianSassine/projectManager/internal/services/jwt-token"

	"github.com/gin-gonic/gin"
)

// Checking if the token is valid on every entry
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedAuthToken)
			return
		}

		if err := jwtToken.ValidateToken(token, false); err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedAuthToken)
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
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToValidateUser)
			return
		}

		tokenInfo, err := jwtToken.ExtractClaims(token)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToValidateUser)
			return
		}

		userId, err := strconv.Atoi(tokenInfo.Subject)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToValidateUser)
			return
		}

		projectId, err := strconv.Atoi(c.Param("projectId"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToValidateUser)
			return
		}

		if !db.IsUserInProject(userId, projectId) {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToValidateUser)
			return
		}

		c.Next()
	}
}
