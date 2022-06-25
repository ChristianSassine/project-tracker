package handlers

import (
	"net/http"

	"github.com/ChristianSassine/projectManager/internal/api"
	logType "github.com/ChristianSassine/projectManager/internal/common/log-type"
	projectErrors "github.com/ChristianSassine/projectManager/internal/common/project-errors"
	"github.com/ChristianSassine/projectManager/internal/services/db"
	"github.com/ChristianSassine/projectManager/internal/services/encryption"

	"github.com/gin-gonic/gin"
)

func GetAllProjects(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchProjects)
			return
		}

		projects, err := db.GetUserProjects(id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchProjects)
			return
		}

		c.JSON(http.StatusOK, projects)
	}
}

func AddProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the necessary information
		requestInfo := api.Project{}
		userId, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddProject)
			return
		}

		if err := c.ShouldBind(&requestInfo); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddProject)
			return
		}

		// Encrypting the project's password (for extra security)
		encryptedPassword, err := encryption.EncryptPassword(requestInfo.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddProject)
			return
		}

		// Adding project to the database
		project, err := db.CreateProject(userId, requestInfo.Title, encryptedPassword)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddProject)
			return
		}
		c.JSON(http.StatusCreated, project)

		// Saving event in the logs
		go db.AddLog(userId, project.Id, logType.ProjectCreated, project.Title)
	}
}

func AddUserToProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the request body
		requestInfo := api.ProjectJoinRequest{}
		if err := c.ShouldBind(&requestInfo); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddUserProject)
			return
		}

		// Verifications
		if !checkProjectPassword(requestInfo.Id, requestInfo.Password, db) {
			// Add error here
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Getting necessary information
		userId, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddUserProject)
			return
		}
		projectTitle, err := db.GetProjectTitle(requestInfo.Id)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToAddUserProject)
			return
		}

		// Adding user to the project
		if err := db.AddUserToProject(userId, requestInfo.Id); err != nil {
			c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedToAddUserProject)
			return
		}
		c.JSON(http.StatusCreated, api.Project{Title: projectTitle, Id: requestInfo.Id})

		// Saving event in the logs
		go db.AddLog(userId, requestInfo.Id, logType.ProjectJoined)
	}
}

func DeleteProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToDeleteProject)
			return
		}

		if err := db.DeleteProject(projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToDeleteProject)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func checkProjectPassword(projectId int, password string, db *db.DB) bool {
	encryptedPassword, err := db.GetProjectPassword(projectId)
	if err != nil {
		return false
	}
	return encryption.CheckPassword(password, encryptedPassword)
}
