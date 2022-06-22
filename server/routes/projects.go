package routes

import (
	"BugTracker/api"
	logType "BugTracker/common"
	"BugTracker/handlers"
	"BugTracker/services/db"
	"BugTracker/services/encryption"
	"BugTracker/utilities"
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
		// TODO: Change it to id, would be better
		projects, err := db.GetUserProjects(tokenClaims.Username)

		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, projects)
	})

	// Project Creation endpoint
	r.POST("/project", func(c *gin.Context) {
		requestInfo := api.Project{}

		userId, err := getUserId(c)
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

		encryptedPassword, err := encryption.EncryptPassword(requestInfo.Password)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		project, err := db.CreateProject(userId, requestInfo.Title, encryptedPassword)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Adding to the logs
		go db.AddLog(userId, project.Id, logType.ProjectCreated, project.Title)

		log.PrintInfo("New project :'", project.Title, "' has been created")

		c.JSON(http.StatusCreated, project)
	})

	// Join project
	r.POST("/project/join", func(c *gin.Context) {
		requestInfo := api.ProjectJoinRequest{}

		if err := c.ShouldBind(&requestInfo); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !checkProjectPassword(requestInfo.Id, requestInfo.Password, db) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, err := getUserId(c)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projectTitle, err := db.GetProjectTitle(requestInfo.Id)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := db.AddUserToProject(userId, requestInfo.Id); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Adding to the logs
		go db.AddLog(userId, requestInfo.Id, logType.ProjectJoined)

		c.JSON(http.StatusCreated, api.Project{Title: projectTitle, Id: requestInfo.Id})
	})

	// Delete project
	// TODO: might need refactor
	r.DELETE("/project/:projectId", handlers.ValidateUserProject(db), func(c *gin.Context) {
		projectId, err := strconv.Atoi(c.Param("projectId"))
		if err != nil {
			utilities.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := db.DeleteProject(projectId); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	})
}

func checkProjectPassword(projectId int, password string, db *db.DB) bool {
	encryptedPassword, err := db.GetProjectPassword(projectId)
	if err != nil {
		log.PrintWarning("HERE2")
		return false
	}
	log.PrintWarning(encryptedPassword)
	log.PrintWarning(password)
	return encryption.CheckPassword(password, encryptedPassword)
}
