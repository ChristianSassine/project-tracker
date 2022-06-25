package routes

import (
	"github.com/ChristianSassine/projectManager/internal/handlers"
	"github.com/ChristianSassine/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(r *gin.RouterGroup, db *db.DB) {

	// Validate a jwt token
	r.GET("/projects", handlers.GetAllProjects(db))

	// Project Creation endpoint
	r.POST("/project", handlers.AddProject(db))

	// Join project endpoint
	r.POST("/project/join", handlers.AddUserToProject(db))

	// Delete project endpoint
	r.DELETE("/project/:projectId", handlers.ValidateUserProject(db), handlers.DeleteProject(db))
}
