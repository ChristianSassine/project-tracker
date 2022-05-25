package routes

import (
	"BugTracker/middlewares"
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(r *gin.RouterGroup, db *db.DB) {
	group := r.Group("/data", middlewares.ValidTokenMiddleware())

	// Validate a jwt token
	group.GET("/projects", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		info, err := jwtToken.ExtractInformation(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projects, err := db.GetUserProjects(info.Username)

		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, projects)
	})
}
