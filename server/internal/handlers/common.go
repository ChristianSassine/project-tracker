package handlers

import (
	"net/http"
	"strconv"

	"github.com/ChristianSassine/projectManager/internal/api"
	log "github.com/ChristianSassine/projectManager/internal/log"
	"github.com/ChristianSassine/projectManager/internal/services/db"
	jwtToken "github.com/ChristianSassine/projectManager/internal/services/jwt-token"

	"github.com/gin-gonic/gin"
)

func getTokenClaims(c *gin.Context) (*api.Claims, error) {
	token, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		log.PrintError(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return &api.Claims{}, err
	}

	tknInfo, err := jwtToken.ExtractClaims(token)
	if err != nil {
		log.PrintError(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return &api.Claims{}, err
	}
	return tknInfo, nil
}

func getProjectId(c *gin.Context) (int, error) {
	projectIdString := c.Param("projectId")

	projectId, err := strconv.Atoi(projectIdString)
	if err != nil {
		return 0, err
	}
	return projectId, nil
}

func logEvent(c *gin.Context, db *db.DB, logType string, args ...string) {
	// LogEvents are launched in concurrency, it's better to not use the context error handling
	tknInfo, err := getTokenClaims(c)
	if err != nil {
		log.PrintError("Failed to add the log of type '", logType, "'. For the following error: ", err)
		return
	}

	projectId, err := getProjectId(c)
	if err != nil {
		log.PrintError("Failed to add the log of type '", logType, "'. For the following error: ", err)
		return
	}

	userId, err := strconv.Atoi(tknInfo.Subject)
	if err != nil {
		log.PrintError("Failed to add the log of type '", logType, "'. For the following error: ", err)
		return
	}

	go db.AddLog(userId, projectId, logType, args...)
}

func getUserId(c *gin.Context) (int, error) {
	tokenClaims, err := getTokenClaims(c)
	if err != nil {
		return 0, err
	}

	userId, err := strconv.Atoi(tokenClaims.Subject)
	if err != nil {
		log.PrintError(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return 0, err
	}
	return userId, nil
}
