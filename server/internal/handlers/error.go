package handlers

import (
	log "github.com/ChristianSassine/projectManager/internal/log"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// For now, we're printing the errors
		for _, err := range c.Errors {
			log.PrintError(err)
		}
	}
}
