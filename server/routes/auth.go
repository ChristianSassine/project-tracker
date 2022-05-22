package routes

import (
	"BugTracker/api"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// TODO : remove and add database
var users = map[string]string{
	"newGame": "plus",
}

func AuthMiddleware(r *gin.RouterGroup) {
	group := r.Group("/auth")
	group.POST("/login", func(c *gin.Context) {
		creds := &api.LoginCreds{}

		if err := c.ShouldBind(creds); err != nil {
			utilities.ErrorLog.Println("An error has occured :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if _, ok := users[creds.Username]; ok == false {
			utilities.InfoLog.Println("Wrong authentication for the user :", creds.Username)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtTkn, err := jwtToken.GenerateToken(creds.Username, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// TODO: generateRefreshToken
		refreshTkn, err := jwtToken.GenerateToken(creds.Username, true)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.SetCookie("JWT_TOKEN", jwtTkn, 30, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", refreshTkn, 30, "/", "localhost", true, true)
	})
	group.GET("/validate", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := &api.Claims{}

		if _, err := jwtToken.ValidateToken(token, claims); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				utilities.ErrorLog.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		utilities.InfoLog.Println("User", claims.Username, "is validated")
		c.AbortWithStatus(http.StatusOK)
	})
}
