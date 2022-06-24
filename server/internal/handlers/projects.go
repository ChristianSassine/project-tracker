package handlers

import (
	"net/http"

	"github.com/krispier/projectManager/internal/api"
	logType "github.com/krispier/projectManager/internal/common"
	"github.com/krispier/projectManager/internal/services/db"
	"github.com/krispier/projectManager/internal/services/encryption"

	"github.com/gin-gonic/gin"
)

func GetAllProjects(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		projects, err := db.GetUserProjects(id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
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
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := c.ShouldBind(&requestInfo); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Encrypting the project's password (for extra security)
		encryptedPassword, err := encryption.EncryptPassword(requestInfo.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Adding project to the database
		project, err := db.CreateProject(userId, requestInfo.Title, encryptedPassword)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusCreated, project)

		// Adding to the logs
		go db.AddLog(userId, project.Id, logType.ProjectCreated, project.Title)
	}
}

func AddUserToProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the request body
		requestInfo := api.ProjectJoinRequest{}
		if err := c.ShouldBind(&requestInfo); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
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
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		projectTitle, err := db.GetProjectTitle(requestInfo.Id)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// Adding user to the project
		if err := db.AddUserToProject(userId, requestInfo.Id); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, api.Project{Title: projectTitle, Id: requestInfo.Id})

		// Adding to the logs
		go db.AddLog(userId, requestInfo.Id, logType.ProjectJoined)
	}
}

func DeleteProject(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := db.DeleteProject(projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
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
