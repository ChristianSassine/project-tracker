package handlers

import (
	"net/http"
	"strconv"

	"github.com/krispier/projectManager/internal/api"
	logType "github.com/krispier/projectManager/internal/common/log-type"
	projectErrors "github.com/krispier/projectManager/internal/common/project-errors"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
			return
		}

		tasks, err := db.GetAllTasks(projectId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, projectErrors.FailedToFetchTasks)
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
				c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
				return
			}

			tasks, err := db.GetTasksByState(projectId, taskState[0])
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
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
				c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
				return
			}
			limit, err := strconv.Atoi(taskLimit[0])
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
				return
			}

			// Getting the tasks
			tasks, err := db.GetTasksWithLimit(projectId, limit)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasks)
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
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasksStats)
			return
		}

		tasksStats, err := db.GetTasksStats(projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchTasksStats)
			return
		}
		c.JSON(http.StatusOK, tasksStats)
	}
}

func AddTask(db *db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddTask)
			return
		}

		task := &api.Task{}
		if err := c.ShouldBind(task); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddTask)
			return
		}

		if err := db.AddTask(task, projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToAddTask)
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
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}

		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToUpdateTask)
			return
		}

		if err := db.UpdateTask(task, projectId); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func DeleteTask(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToDeleteTask)
			return
		}

		taskIdString, ok := c.Request.URL.Query()["id"]
		if !ok {
			// TODO : Add error here
			c.AbortWithError(http.StatusNotFound, projectErrors.FailedToDeleteTask)
			return
		}
		taskId, err := strconv.Atoi(taskIdString[0])
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToDeleteTask)
			return
		}

		taskTitle, err := db.DeleteTask(taskId, projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToDeleteTask)
			return
		}
		c.Status(http.StatusNoContent)

		// Adding to the logs
		go logEvent(c.Copy(), db, logType.TaskDeleted, taskTitle)
	}
}

func UpdateTaskPosition(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		taskPositionRequest := &api.TaskPatchRequest{}
		err := c.ShouldBind(taskPositionRequest)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}

		if err := db.UpdateTaskPosition(taskPositionRequest.PreviousIndex, taskPositionRequest.CurrentIndex, taskPositionRequest.TaskId); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}

		c.Status(http.StatusOK)
	}
}

func UpdateTaskState(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, err := getProjectId(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToUpdateTask)
			return
		}

		taskStateRequest := &api.TaskPatchRequest{}
		err = c.ShouldBind(taskStateRequest)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}

		previousTask, err := db.UpdateTaskState(taskStateRequest.NewState, taskStateRequest.CurrentIndex, taskStateRequest.TaskId, projectId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToUpdateTask)
			return
		}
		c.Status(http.StatusOK)

		// Adding to the logs
		go logEvent(c.Copy(), db, logType.TaskStateModification, previousTask.Title, previousTask.State, taskStateRequest.NewState)
	}
}
