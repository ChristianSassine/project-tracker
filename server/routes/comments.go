package routes

import (
	"BugTracker/api"
	"BugTracker/middlewares"
	"BugTracker/services/db"
	log "BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskCommentsRoutes(r *gin.RouterGroup, db *db.DB) {
	taskGroup := r.Group("", middlewares.ValidUserProjectAccessMiddleware(db))

	// Get a task's comments endpoint
	taskGroup.GET("/project/:projectId/task/:taskId/comments", func(c *gin.Context) {
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		comments, err := db.GetTaskComments(taskId)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, comments)
	})

	// Add a comment to a task endpoint
	taskGroup.POST("/project/:projectId/task/:taskId/comment", func(c *gin.Context) {
		userId, err := getUserId(c)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var request api.CommentAddRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		log.PrintWarning(request)

		if err := db.AddTaskComment(userId, request.Content, taskId); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusCreated)
	})
}
