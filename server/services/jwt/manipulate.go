package jwtToken

import (
	"BugTracker/api"

	"github.com/golang-jwt/jwt"
)

func extractToken(token string) (*api.Claims, error) {

	claims := &api.Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return claims, err
}
