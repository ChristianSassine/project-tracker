package middlewares

import (
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Checking if the token is valid on every entry
func ValidTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := jwtToken.ValidateToken(token, false); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				utilities.ErrorLog.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

// Validating that user has permissions to access the project's tasks
// TODO: maybe needs to be renamed, too long
func ValidUserProjectAccessMiddleware(db *db.DB) gin.HandlerFunc {
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
