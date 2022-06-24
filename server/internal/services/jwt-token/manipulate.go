package jwtToken

import (
	"github.com/krispier/projectManager/internal/api"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractClaims(token string) (*api.Claims, error) {

	claims := &api.Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return claims, err
}
