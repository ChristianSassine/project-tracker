package api

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	Type     string `json:"typ"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
