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

	// Creating a big route
	superGroup := router.Group("/api")

	// Initializing database
	database := &db.DB{}
	database.Connect()

	// Adding routes
	routes.AuthRoutes(superGroup, database)
	routes.ProjectsRoutes(superGroup, database)

	// Launching the server
	router.Run("localhost:8080")
}
