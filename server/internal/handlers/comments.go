package handlers

import (
	"net/http"
	"strconv"

	"github.com/krispier/projectManager/internal/api"
	projectErrors "github.com/krispier/projectManager/internal/common/project-errors"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func AddComment(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		userId, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddComment)
			return
		}

		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddComment)
			return
		}

		var request api.CommentAddRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddComment)
			return
		}

		if err := db.AddTaskComment(userId, request.Content, taskId); err != nil {
			c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedToAddComment)
			return
		}

		c.Status(http.StatusCreated)
	}
}
func GetAllComments(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchComments)
			return
		}

		comments, err := db.GetTaskComments(taskId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, projectErrors.FailedToFetchComments)
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}
