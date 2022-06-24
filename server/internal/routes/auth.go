package routes

import (
	"github.com/krispier/projectManager/internal/handlers"
	"github.com/krispier/projectManager/internal/services/db"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, db *db.DB) {
	group := r.Group("/auth")

	// Login endpoint.
	group.POST("/login", handlers.Login(db))

	// User creation endpoint.
	group.POST("/create", handlers.CreateAccount(db))

	// Fetch username endpoint.
	group.GET("/user", handlers.FetchUsername)

	// Validate a jwt token endpoint.
	group.GET("/validate", handlers.ValidateTkn)

	// Refreshing access token endpoint.
	group.GET("/refresh")

	// Logout endpoint.
	group.GET("/logout", handlers.Logout)
}
