package routes

import (
	"BugTracker/middlewares"
	"BugTracker/services/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(r *gin.RouterGroup, db *db.DB) {
	group := r.Group("/data", middlewares.ValidTokenMiddleware())

	// Validate a jwt token
	group.GET("/projects", func(c *gin.Context) {
		// token, err := c.Cookie("JWT_TOKEN")
		// if err != nil {
		// 	utilities.ErrorLog.Println(err)
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// if err := jwtToken.ValidateToken(token); err != nil {
		// 	if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
		// 		utilities.ErrorLog.Println(err)
		// 		c.AbortWithStatus(http.StatusUnauthorized)
		// 		return
		// 	}
		// 	utilities.ErrorLog.Println(err)
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// utilities.InfoLog.Println("User", claims.Username, "is validated")
		c.AbortWithStatus(http.StatusOK)
	})
}
