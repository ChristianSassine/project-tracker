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

	// Creating a big route
	superGroup := router.Group("/api")
	dataGroup := superGroup.Group("/data", middlewares.ValidTokenMiddleware(), middlewares.ValidUserProjectAccessMiddleware(database))

	// Adding routes
	routes.AuthRoutes(superGroup, database)
	routes.ProjectsRoutes(dataGroup, database)
	routes.TasksRoutes(dataGroup, database)

	// Launching the server
	router.Run("localhost:8080")
}
