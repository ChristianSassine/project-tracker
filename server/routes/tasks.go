package routes

import (
	"BugTracker/api"
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TasksRoutes(r *gin.RouterGroup, db *db.DB) {

	// Getting all tasks handler
	r.GET("/project/:projectId/tasks", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")

		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenInfo, err := jwtToken.ExtractInformation(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(tokenInfo.Subject)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projectIdString := c.Param("projectId")

		projectId, err := strconv.Atoi(projectIdString)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		tasks, err := db.GetTasks(userId, projectId)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusCreated, tasks)
	})

	r.POST("/project/:projectId/task", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")

		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenInfo, err := jwtToken.ExtractInformation(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(tokenInfo.Subject)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projectIdString := c.Param("projectId")

		projectId, err := strconv.Atoi(projectIdString)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// TODO : Maybe create a middleWare for all tasks routes
		if !db.ValidateUserProjectPermission(userId, projectId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		task := &api.Task{}

		if err := c.ShouldBind(task); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := db.AddTask(task, projectId); err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusCreated)
	})
}
