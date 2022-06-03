package routes

import (
	"BugTracker/api"
	"BugTracker/services/db"
	jwtToken "BugTracker/services/jwt"
	"BugTracker/utilities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRoutes(r *gin.RouterGroup, db *db.DB) {
	group := r.Group("/auth")

	// Login with identifiers
	group.POST("/login", func(c *gin.Context) {
		creds := &api.LoginCreds{}

		if err := c.ShouldBind(creds); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		id, valid := db.ValidateUser(creds.Username, creds.Password)
		if !valid {
			utilities.InfoLog.Println("Wrong authentication for the user :", creds.Username)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtTkn, err := jwtToken.GenerateToken(creds.Username, id, time.Minute*60, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		refreshTkn, err := jwtToken.GenerateToken(creds.Username, id, time.Hour*24, true)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		utilities.InfoLog.Print("access token:", jwtTkn)
		utilities.InfoLog.Print("refresh token:", refreshTkn)

		c.SetCookie("JWT_TOKEN", jwtTkn, 60*60, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", refreshTkn, 4*60, "/auth/refresh", "localhost", true, true)
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

		id, err := db.AddUser(creds.Username, creds.Password, creds.Email)
		if err != nil {
			utilities.InfoLog.Println("User", creds.Username, "already exists || Error :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		utilities.InfoLog.Println("User", creds.Username, "has been created")
		c.Status(http.StatusCreated)

		jwtTkn, err := jwtToken.GenerateToken(creds.Username, id, time.Second*15, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		refreshTkn, err := jwtToken.GenerateToken(creds.Username, id, time.Minute*10, true)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.SetCookie("JWT_TOKEN", jwtTkn, 60*60, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", refreshTkn, 60*10, "/auth/refresh", "localhost", true, true)
	})

	// Validate a jwt token
	group.GET("/validate", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		utilities.InfoLog.Print("Validation token:", token)

		if err := jwtToken.ValidateToken(token, false); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				utilities.ErrorLog.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		utilities.InfoLog.Println("User", claims.Username, "is validated")
		c.JSON(http.StatusOK, claims.Username)
	})

	// TODO : Handle refresh token
	group.GET("/refresh", func(c *gin.Context) {
		token, err := c.Cookie("JWT_REFRESH")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := jwtToken.ValidateToken(token, true); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				utilities.ErrorLog.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, _ := strconv.Atoi(claims.Subject)

		jwtTkn, err := jwtToken.GenerateToken(claims.Username, id, time.Minute*5, false)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.SetCookie("JWT_TOKEN", jwtTkn, 60*5, "/", "localhost", true, true)
		utilities.InfoLog.Println("User", claims.Username, "has refreshed his token")
		c.JSON(http.StatusOK, claims.Username)
	})

	group.GET("/logout", func(c *gin.Context) {
		jwtTknOverwrite := "loggedOut"
		jwtRefreshOverwrite := "loggedOut"

		c.Status(http.StatusOK)
		c.SetCookie("JWT_TOKEN", jwtTknOverwrite, 1, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", jwtRefreshOverwrite, 1, "/auth/refresh", "localhost", true, true)
	})
}
