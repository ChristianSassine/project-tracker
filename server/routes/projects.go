package routes

import (
	"BugTracker/api"
	logType "BugTracker/common"
	"BugTracker/services/db"
	log "BugTracker/utilities"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(r *gin.RouterGroup, db *db.DB) {

	// Validate a jwt token
	r.GET("/projects", func(c *gin.Context) {
		tokenClaims, err := getTokenClaims(c)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projects, err := db.GetUserProjects(tokenClaims.Username)

		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, projects)
	})

	// Project Creation handler
	r.POST("/project", func(c *gin.Context) {
		requestInfo := api.Project{}

		tokenClaims, err := getTokenClaims(c)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := c.ShouldBind(&requestInfo); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(tokenClaims.Subject)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		project, err := db.CreateProject(id, requestInfo.Title)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Adding to the logs
		go db.AddLog(id, project.Id, logType.ProjectCreation, project.Title)

		log.PrintInfo("New project :'", project.Title, "' has been created")

		c.JSON(http.StatusCreated, project)
	})
}
