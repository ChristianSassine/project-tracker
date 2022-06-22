package handlers

import (
	"BugTracker/api"
	"BugTracker/services/db"
	"BugTracker/services/encryption"
	jwtToken "BugTracker/services/jwt-token"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	secondsToMinute, minuteToHour                = 60, 60
	jwtCookieExpiration, refreshCookieExpiration = secondsToMinute * minuteToHour, 0
	tknExpiration, refreshExpiration             = time.Minute * 60, time.Hour * 24
)

func ValidateTkn(c *gin.Context) {
	// Fetch token
	token, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Validation and send the result
	if err := jwtToken.ValidateToken(token, false); err != nil {
		if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusOK)
}

func RefreshTkn(c *gin.Context) {
	// Extract the refresh token from the cookies
	token, err := c.Cookie("JWT_REFRESH")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Validate the token
	if err := jwtToken.ValidateToken(token, true); err != nil {
		if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Extracting claims from them
	claims, err := jwtToken.ExtractClaims(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Generating token
	id, _ := strconv.Atoi(claims.Subject)
	jwtTkn, err := jwtToken.GenerateToken(claims.Username, id, tknExpiration, false)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Setting token as cookie and sending it
	c.SetCookie("JWT_TOKEN", jwtTkn, secondsToMinute*minuteToHour, "/", "localhost", true, true)
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	jwtTknOverwrite, jwtRefreshOverwrite := "loggedOut", "loggedOut"

	// Overwriting tokens with expiring tokens
	c.Status(http.StatusOK)
	c.SetCookie("JWT_TOKEN", jwtTknOverwrite, 1, "/", "localhost", true, true)
	c.SetCookie("JWT_REFRESH", jwtRefreshOverwrite, 1, "/api/auth/refresh", "localhost", true, true)
}

func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetching the data from the body
		creds := &api.LoginCreds{}
		if err := c.ShouldBind(creds); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// Verifying if the user has the valid information
		id, err := db.GetUserId(creds.Username)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		if !checkUserPassword(id, creds.Password, db) {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// Sending him auth tokens to access the application
		if err := setTokens(c, creds.Username, id); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)
	}
}

func CreateAccount(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetching the data from the body
		creds := &api.RegistrationCreds{}
		if err := c.ShouldBind(creds); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// TODO : remove handling from controller
		// Checking if the user actually exists
		if ok, err := db.CheckIfUserExists(creds.Username); ok {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// Encrypting the password
		encryptedPassword, err := encryption.EncryptPassword(creds.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Adding the user with the encrypted password in the database
		if err := db.AddUser(creds.Username, encryptedPassword, creds.Email); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Status(http.StatusCreated)
	}
}

func FetchUsername(c *gin.Context) {
	// Extracting the name from the token (We know it wasn't tempered because of the verification middleware)
	token, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	claims, err := jwtToken.ExtractClaims(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Sending the username to the client
	c.JSON(http.StatusOK, claims.Username)
}

func setTokens(c *gin.Context, username string, id int) error {
	// Generating the tokens
	jwtTkn, err := jwtToken.GenerateToken(username, id, tknExpiration, false)
	if err != nil {
		return err
	}
	refreshTkn, err := jwtToken.GenerateToken(username, id, refreshExpiration, true)
	if err != nil {
		return err
	}

	// Setting the tokens
	// TODO: set the origin automatically
	c.SetCookie("JWT_TOKEN", jwtTkn, jwtCookieExpiration, "/", "localhost", true, true)
	c.SetCookie("JWT_REFRESH", refreshTkn, refreshCookieExpiration, "/api/auth/refresh", "localhost", true, true)

	return nil
}

func checkUserPassword(userId int, password string, db *db.DB) bool {
	encryptedPassword, err := db.GetUserPassword(userId)
	if err != nil {
		return false
	}
	return encryption.CheckPassword(password, encryptedPassword)
}
