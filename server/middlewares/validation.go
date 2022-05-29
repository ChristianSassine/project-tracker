package middlewares

import (
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Checking if the token is valid on every entry
func ValidTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
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
