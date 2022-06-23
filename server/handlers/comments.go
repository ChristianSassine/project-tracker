package handlers

import (
	"BugTracker/api"
	"BugTracker/services/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		userId, err := getUserId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var request api.CommentAddRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := db.AddTaskComment(userId, request.Content, taskId); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusCreated)
	}
}
func GetAllComments(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		comments, err := db.GetTaskComments(taskId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}
