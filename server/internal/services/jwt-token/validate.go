package jwtToken

import (
	"errors"

	"github.com/ChristianSassine/projectManager/internal/api"

	"github.com/golang-jwt/jwt/v4"
)

var UnvalidTokenError error = errors.New("The token is not valid")

func ValidateToken(token string, isRefreshToken bool) error {

	claims := &api.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return err
	}

	if !tkn.Valid {
		return UnvalidTokenError
	}

	if isRefreshToken && claims.Type != "refresh" {
		return UnvalidTokenError
	}

	if !isRefreshToken && claims.Type != "validation" {
		err = UnvalidTokenError
	}

	return err
}
