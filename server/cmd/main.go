package main

import (
	"BugTracker/routes"
	"BugTracker/services/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO : Needs to be remplaced in case of another client address
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	// Initializing the server and middlewares
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Creating a big route
	superGroup := router.Group("/api")

	// Initializing database
	database := &db.DB{}
	database.Connect()

	// Adding routes
	routes.AuthMiddleware(superGroup, database)

	// Launching the server
	router.Run("localhost:8080")
}
