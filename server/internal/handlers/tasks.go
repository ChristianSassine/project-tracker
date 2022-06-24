package handlers

import (
	"net/http"
	"strconv"

	"github.com/krispier/projectManager/internal/api"
	logType "github.com/krispier/projectManager/internal/common"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		tasks, err := db.GetAllTasks(projectId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func GetTasksByState(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskState, ok := c.Request.URL.Query()["state"]
		if ok {
			projectId, err := getProjectId(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			tasks, err := db.GetTasksByState(projectId, taskState[0])
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusOK, tasks)
		}
	}
}

func GetTasksByLimit(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskLimit, ok := c.Request.URL.Query()["limit"]
		if ok {
			// Extracting necessary info
			projectId, err := getProjectId(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			limit, err := strconv.Atoi(taskLimit[0])
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			// Getting the tasks
			tasks, err := db.GetTasksWithLimit(projectId, limit)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusOK, tasks)
			return
		}
	}
}

func GetTaskStats(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		tasksStats, err := db.GetTasksStats(projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, tasksStats)
	}
}

func AddTask(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		task := &api.Task{}
		if err := c.ShouldBind(task); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := db.AddTask(task, projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusCreated)

		// Adding to the logs
		go logEvent(c.Copy(), db, logType.TaskCreation, task.Title)
	}
}

func UpdateTask(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		task := &api.Task{}

		if err := c.ShouldBind(task); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if err := db.UpdateTask(task, projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func DeleteTask(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		taskIdString, ok := c.Request.URL.Query()["id"]
		if !ok {
			// TODO : Add error here
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		taskId, err := strconv.Atoi(taskIdString[0])
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		taskTitle, err := db.DeleteTask(taskId, projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusNoContent)

		// Adding to the logs
		go logEvent(c.Copy(), db, logType.TaskDeleted, taskTitle)
	}
}

func UpdateTaskPosition(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		taskPositionRequest := &api.TaskPatchRequest{}
		err = c.ShouldBind(taskPositionRequest)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := db.UpdateTaskPosition(taskPositionRequest.PreviousIndex, taskPositionRequest.CurrentIndex, taskPositionRequest.TaskId, projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func UpdateTaskState(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		taskStateRequest := &api.TaskPatchRequest{}
		err = c.ShouldBind(taskStateRequest)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		previousTask, err := db.UpdateTaskState(taskStateRequest.NewState, taskStateRequest.CurrentIndex, taskStateRequest.TaskId, projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusOK)

		// Adding to the logs
		go logEvent(c.Copy(), db, logType.TaskStateModification, previousTask.Title, previousTask.State, taskStateRequest.NewState)
	}
}
