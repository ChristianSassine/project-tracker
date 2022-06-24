package main

import (
	"github.com/krispier/projectManager/internal/handlers"
	"github.com/krispier/projectManager/internal/routes"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing the server and middlewares
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(handlers.ErrorHandler(), handlers.CORSMiddleware())
	router.SetTrustedProxies(nil)

	// Initializing database
	database := &db.DB{}
	database.Connect()

	// Creating the route groups
	superGroup := router.Group("/api")
	dataGroup := superGroup.Group("", handlers.ValidateToken())

	// Adding the routes to the groups
	routes.AuthRoutes(superGroup, database)
	routes.ProjectsRoutes(dataGroup, database)
	routes.TasksRoutes(dataGroup, database)
	routes.TaskCommentsRoutes(dataGroup, database)
	routes.ProjectLogsRoutes(dataGroup, database)

	// Launching the server
	router.Run("localhost:8080")
}
