package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/krispier/projectManager/internal/api"
	projectErrors "github.com/krispier/projectManager/internal/common/project-errors"
	"github.com/krispier/projectManager/internal/services/db"
	"github.com/krispier/projectManager/internal/services/encryption"
	jwtToken "github.com/krispier/projectManager/internal/services/jwt-token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
		c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedAuthToken)
		return
	}

	// Validation and send the result
	if err := jwtToken.ValidateToken(token, false); err != nil {
		if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedAuthToken)
			return
		}
		c.AbortWithError(http.StatusBadRequest, projectErrors.FailedAuthToken)
		return
	}
	c.Status(http.StatusOK)
}

func RefreshTkn(c *gin.Context) {
	// Extract the refresh token from the cookies
	token, err := c.Cookie("JWT_REFRESH")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedRefreshToken)
		return
	}

	// Validate the token
	if err := jwtToken.ValidateToken(token, true); err != nil {
		if err == jwt.ErrSignatureInvalid || err == jwtToken.UnvalidTokenError {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedRefreshToken)
			return
		}
		c.AbortWithError(http.StatusBadRequest, projectErrors.FailedRefreshToken)
		return
	}

	// Extracting claims from them
	claims, err := jwtToken.ExtractClaims(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, projectErrors.FailedRefreshToken)
		return
	}

	// Generating token
	id, _ := strconv.Atoi(claims.Subject)
	jwtTkn, err := jwtToken.GenerateToken(claims.Username, id, tknExpiration, false)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedRefreshToken)
		return
	}

	// Setting token as cookie and sending it
	domain, err := getDomain(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedRefreshToken)
		return
	}
	c.SetCookie("JWT_TOKEN", jwtTkn, secondsToMinute*minuteToHour, "/", domain, true, true)
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	jwtTknOverwrite, jwtRefreshOverwrite := "loggedOut", "loggedOut"

	domain, err := getDomain(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedToLogout)
		return
	}
	// Overwriting tokens with expiring tokens
	c.Status(http.StatusOK)
	c.SetCookie("JWT_TOKEN", jwtTknOverwrite, 1, "/", domain, true, true)
	c.SetCookie("JWT_REFRESH", jwtRefreshOverwrite, 1, "/api/auth/refresh", domain, true, true)
}

func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetching the data from the body
		creds := &api.LoginCreds{}
		if err := c.ShouldBind(creds); err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToLogin)
			return
		}

		// Verifying if the user has the valid information
		id, err := db.GetUserId(creds.Username)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToLogin)
			return
		}
		if !checkUserPassword(id, creds.Password, db) {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToLogin)
			return
		}

		// Sending him auth tokens to access the application
		if err := setTokens(c, creds.Username, id); err != nil {
			c.AbortWithError(http.StatusInternalServerError, projectErrors.FailedToLogin)
			return
		}
		c.Status(http.StatusOK)
	}
}

func IsLoggedOut(c *gin.Context) {
	// If the cookies don't exist or have a problem, the user is disconnected
	jwtTkn, tknErr := c.Cookie("JWT_TOKEN")
	refreshTkn, refreshErr := c.Cookie("JWT_REFRESH")
	if tknErr != nil && refreshErr != nil {
		c.Status(http.StatusOK)
		return
	}

	tknErr = jwtToken.ValidateToken(jwtTkn, false)
	refreshErr = jwtToken.ValidateToken(refreshTkn, true)
	if tknErr != nil && refreshErr != nil {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusUnauthorized)
}

func CreateAccount(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetching the data from the body
		creds := &api.RegistrationCreds{}
		if err := c.ShouldBind(creds); err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToCreateUser)
			return
		}

		// TODO : remove handling from controller
		// Checking if the user actually exists
		if ok, err := db.CheckIfUserExists(creds.Username); ok || err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToCreateUser)
			return
		}

		// Encrypting the password
		encryptedPassword, err := encryption.EncryptPassword(creds.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToCreateUser)
			return
		}

		// Adding the user with the encrypted password in the database
		if err := db.AddUser(creds.Username, encryptedPassword, creds.Email); err != nil {
			c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToCreateUser)
			return
		}

		c.Status(http.StatusCreated)
	}
}

func FetchUsername(c *gin.Context) {
	// Extracting the name from the token (We know it wasn't tempered because of the verification middleware)
	token, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, projectErrors.FailedToFetchUsername)
		return
	}
	claims, err := jwtToken.ExtractClaims(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, projectErrors.FailedToFetchUsername)
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

	domain, err := getDomain(c)
	if err != nil {
		return err
	}

	// Setting the tokens
	c.SetCookie("JWT_TOKEN", jwtTkn, jwtCookieExpiration, "/", domain, true, true)
	c.SetCookie("JWT_REFRESH", refreshTkn, refreshCookieExpiration, "/api/auth/refresh", domain, true, true)

	return nil
}

func checkUserPassword(userId int, password string, db *db.DB) bool {
	encryptedPassword, err := db.GetUserPassword(userId)
	if err != nil {
		return false
	}
	return encryption.CheckPassword(password, encryptedPassword)
}

func getDomain(c *gin.Context) (string, error) {
	parsedUrl, err := url.Parse(c.Request.Header.Get("Origin"))
	if err != nil {
		return "", err
	}
	return parsedUrl.Hostname(), nil
}
