package main

import (
	"github.com/ChristianSassine/projectManager/internal/handlers"
	"github.com/ChristianSassine/projectManager/internal/routes"
	"github.com/ChristianSassine/projectManager/internal/services/db"

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
	subGroup := superGroup.Group("", handlers.ValidateToken())

	// Adding the routes to the groups
	routes.AuthRoutes(superGroup, database)
	routes.ProjectsRoutes(subGroup, database)
	routes.TasksRoutes(subGroup, database)
	routes.TaskCommentsRoutes(subGroup, database)
	routes.ProjectLogsRoutes(subGroup, database)

	// Launching the server
	router.Run()
}
