package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var users = map[string]string{
	"newGame": "plus",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var signingKey []byte = []byte("davidLavariete")

func signing() string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: "HelloMister",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}).SignedString([]byte(signingKey))

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	// fmt.Println("the following token has been created :", token)
	log.Println("the following token has been created :", token)
	return token
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	router := gin.Default()
	router.Use(CORSMiddleware())
	superGroup := router.Group("/api")
	superGroup.POST("", func(c *gin.Context) {
		creds := &Credentials{}

		if err := c.ShouldBind(creds); err != nil {
			log.Println("An error has occured :", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if _, ok := users[creds.Username]; ok == false {
			log.Println("Wrong authentication for the user :", creds.Username)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.SetCookie("JWT_TOKEN", signing(), 30, "/", "localhost", true, true)
		c.SetCookie("JWT_REFRESH", signing(), 30, "/", "localhost", true, true)
	})
	superGroup.GET("/auth", func(c *gin.Context) {
		token, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if !tkn.Valid {
			log.Println("Token not valid :", tkn)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		log.Println("User", claims.Username, "successfully authorized")
		c.AbortWithStatus(http.StatusOK)
	})
	router.Run("localhost:8080")
}
