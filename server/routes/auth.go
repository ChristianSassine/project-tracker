package routes

import (
	"BugTracker/api"
	"BugTracker/services/db"
	"BugTracker/services/encryption"
	jwtToken "BugTracker/services/jwt"
	log "BugTracker/utilities"
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
			log.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id, err := db.GetUserId(creds.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if !checkUserPassword(id, creds.Password, db) {
			log.PrintInfo("Wrong authentication for the user :", creds.Username)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := setTokens(c, creds.Username, id); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	})

	// Create a user endpoint.
	group.POST("/create", func(c *gin.Context) {
		creds := &api.RegistrationCreds{}

		if err := c.ShouldBind(creds); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if exists, err := db.CheckIfUserExists(creds.Username); exists {
			log.PrintInfo("User", creds.Username, "already exists || Error :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		encryptedPassword, err := encryption.EncryptPassword(creds.Password)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, err := db.AddUser(creds.Username, encryptedPassword, creds.Email)
		if err != nil {
			log.PrintInfo("User", creds.Username, "already exists || Error :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := setTokens(c, creds.Username, id); err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		// TODO: remove useless logs
		log.PrintInfo("User", creds.Username, "has been created")

		c.Status(http.StatusCreated)
	})

	// Fetch username endpoint.
	group.GET("/user", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, claims.Username)
	})

	// Validate a jwt token endpoint.
	group.GET("/validate", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := jwtToken.ValidateToken(token, false); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				log.PrintError(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// TODO : Might need to remove this piece of code
		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		log.PrintInfo("User", claims.Username, "has validated his token")
		c.Status(http.StatusOK)
	})

	// TODO : Handle refresh token + fix time of expiration
	// Refreshing access token endpoint.
	group.GET("/refresh", func(c *gin.Context) {
		token, err := c.Cookie("JWT_REFRESH")
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := jwtToken.ValidateToken(token, true); err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
				log.PrintError(err)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, err := jwtToken.ExtractClaims(token)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, _ := strconv.Atoi(claims.Subject)

		jwtTkn, err := jwtToken.GenerateToken(claims.Username, id, time.Minute*60, false)
		if err != nil {
			log.PrintError(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		secondsToMinute, minuteToHour := 60, 60
		c.SetCookie("JWT_TOKEN", jwtTkn, secondsToMinute*minuteToHour, "/", "localhost", true, true)
		log.PrintInfo("User", claims.Username, "has refreshed his token")
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
	// TODO : Remove imaginary numbers
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

func checkUserPassword(userId int, password string, db *db.DB) bool {
	encryptedPassword, err := db.GetUserPassword(userId)
	if err != nil {
		return false
	}
	return encryption.CheckPassword(password, encryptedPassword)
}
