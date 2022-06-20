package main

import (
	"BugTracker/middlewares"
	"BugTracker/routes"
	"BugTracker/services/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing the server and middlewares
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.SetTrustedProxies(nil)

	// Initializing database
	database := &db.DB{}
	database.Connect()

	// Creating the route groups
	superGroup := router.Group("/api")
	dataGroup := superGroup.Group("/data", middlewares.ValidTokenMiddleware())

	// Adding the routes to the groups
	routes.AuthRoutes(superGroup, database)
	routes.ProjectsRoutes(dataGroup, database)
	routes.TasksRoutes(dataGroup, database)
	routes.TaskCommentsRoutes(dataGroup, database)
	routes.ProjectLogsRoutes(dataGroup, database)

	// Launching the server
	router.Run("localhost:8080")
}
