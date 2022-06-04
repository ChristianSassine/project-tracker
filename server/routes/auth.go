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

		if err := setTokens(c, creds.Username, id); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusOK)
	})

	// Create a user endpoint.
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

		if err := setTokens(c, creds.Username, id); err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		// TODO: remove useless logs
		utilities.InfoLog.Println("User", creds.Username, "has been created")

		c.Status(http.StatusCreated)
	})

	// Fetch username endpoint.
	group.GET("/user", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO : Might need to remove this piece of code
		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		utilities.InfoLog.Println("User", claims.Username, "has validated his token")
		c.Status(http.StatusOK)
	})

	// Validate a jwt token endpoint.
	group.GET("/validate", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

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

		// TODO : Might need to remove this piece of code
		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			utilities.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		utilities.InfoLog.Println("User", claims.Username, "has validated his token")
		c.Status(http.StatusOK)
	})

	// TODO : Handle refresh token
	// Refreshing access token endpoint.
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

	// Logout endpoint.
	group.GET("/logout", func(c *gin.Context) {
		jwtTknOverwrite, jwtRefreshOverwrite := "loggedOut", "loggedOut"

		c.Status(http.StatusOK)
		c.SetCookie("JWT_TOKEN", jwtTknOverwrite, 1, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", jwtRefreshOverwrite, 1, "/api/auth/refresh", "localhost", true, true)
	})
}

func setTokens(c *gin.Context, username string, id int) error {
	secondsToMinute, minuteToHour := 60, 60

	jwtTkn, err := jwtToken.GenerateToken(username, id, time.Minute*60, false)
	if err != nil {
		return err
	}

	refreshTkn, err := jwtToken.GenerateToken(username, id, time.Hour*24, true)
	if err != nil {
		return err
	}

	c.SetCookie("JWT_TOKEN", jwtTkn, secondsToMinute*minuteToHour, "/", "localhost", true, true)
	c.SetCookie("JWT_REFRESH", refreshTkn, 0, "/api/auth/refresh", "localhost", true, true)

	return nil
}
