package routes

import (
	"BugTracker/api"
	"BugTracker/middlewares"
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"
	"strconv"

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

		tokenInfo, err := jwtToken.ExtractInformation(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projects, err := db.GetUserProjects(tokenInfo.Username)

		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, projects)
	})

	// Getting all tasks handler
	group.GET("/project/:projectId/tasks", func(c *gin.Context) {
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

	// Project Creating handler
	group.POST("/project", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		requestInfo := api.Project{}

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

		if err := c.ShouldBind(&requestInfo); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(info.Subject)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := db.CreateProject(id, requestInfo.Title); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		utilities.InfoLog.Print("New project :'", requestInfo.Title, "' has been created")
		c.AbortWithStatus(http.StatusCreated)
	})
}
