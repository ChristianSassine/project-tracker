package routes

import (
	"BugTracker/api"
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(r *gin.RouterGroup, db *db.DB) {
	group := r.Group("/auth")

	// Login with identifiers
	group.POST("/login", func(c *gin.Context) {
		creds := &api.LoginCreds{}

		if err := c.ShouldBind(creds); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if valid := db.ValidateUser(creds.Username, creds.Password); !valid {
			utilities.InfoLog.Println("Wrong authentication for the user :", creds.Username)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtTkn, err := jwtToken.GenerateToken(creds.Username, time.Minute*5, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// TODO: handle Refresh token
		refreshTkn, err := jwtToken.GenerateToken(creds.Username, time.Minute*10, true)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.SetCookie("JWT_TOKEN", jwtTkn, 60*5, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", refreshTkn, 60*10, "/", "localhost", true, true)
	})

	// Validate a jwt token
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

	// Create a user
	group.POST("/create", func(c *gin.Context) {
		creds := &api.RegistrationCreds{}

		if err := c.ShouldBind(creds); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if exists, err := db.CheckIfUserExists(creds.Username); exists {
			utilities.InfoLog.Println("User", creds.Username, "already exists || Error :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err := db.AddUser(creds.Username, creds.Password, creds.Email); err != nil {
			utilities.InfoLog.Println("User", creds.Username, "already exists || Error :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		utilities.InfoLog.Println("User", creds.Username, "has been created")
		c.Status(http.StatusCreated)

		jwtTkn, err := jwtToken.GenerateToken(creds.Username, time.Minute*5, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// TODO: handle Refresh token
		refreshTkn, err := jwtToken.GenerateToken(creds.Username, time.Minute*10, true)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.SetCookie("JWT_TOKEN", jwtTkn, 60*5, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", refreshTkn, 60*10, "/", "localhost", true, true)
	})

	// TODO : Handle refresh token
}
