package routes

import (
	"BugTracker/api"
	"BugTracker/middlewares"
	"BugTracker/services/db"
	"BugTracker/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TasksRoutes(r *gin.RouterGroup, db *db.DB) {
	taskGroup := r.Group("", middlewares.ValidUserProjectAccessMiddleware(db))

	// Getting all tasks endpoint
	taskGroup.GET("/project/:projectId/tasks", func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		urlQueries := c.Request.URL.Query()
		taskState, ok := urlQueries["State"]
		if !ok {
			tasks, err := db.GetAllTasks(projectId)
			if err != nil {
				utilities.ErrorLog.Println(err)
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			c.JSON(http.StatusCreated, tasks)
			return
		}

		tasks, err := db.GetTasksByState(projectId, taskState[0])
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusCreated, tasks)
	})

	// Creating a task endpoint
	taskGroup.POST("/project/:projectId/task", func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
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

	// Updating a task endpoint
	taskGroup.PUT("/project/:projectId/task", func(c *gin.Context) {
		task := &api.Task{}

		if err := c.ShouldBind(task); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		projectId, err := getProjectId(c)
		if err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := db.UpdateTask(task, projectId); err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusCreated)
	})

	// Deleting a task endpoint
	taskGroup.DELETE("/project/:projectId/task", func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		urlQueries := c.Request.URL.Query()
		taskIdString, ok := urlQueries["id"]
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		taskId, err := strconv.Atoi(taskIdString[0])
		if err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := db.DeleteTask(taskId, projectId); err != nil {
			utilities.ErrorLog.Print(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusCreated)
	})
}

func getProjectId(c *gin.Context) (int, error) {
	projectIdString := c.Param("projectId")

	projectId, err := strconv.Atoi(projectIdString)
	if err != nil {
		return 0, err
	}
	return projectId, nil
}
