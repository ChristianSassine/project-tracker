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
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func signing() string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: "HelloMister",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 15).Unix(),
		},
	}).SignedString([]byte("davidLavariete"))

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
	log.SetPrefix("[LOG] -> ")
	log.SetFlags(log.Lmsgprefix | log.LstdFlags)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/", func(c *gin.Context) {
		c.SetCookie("JWT_TOKEN", signing(), 15, "/", "localhost", true, false)
		c.SetCookie("JWT_REFRESH", signing(), 15, "/", "localhost", true, false)
	})

	router.GET("/", func(c *gin.Context) {
		_, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Status(http.StatusOK)
	})

	router.Run("localhost:8080")
}
