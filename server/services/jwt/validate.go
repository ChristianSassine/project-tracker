package jwtToken

import (
	"BugTracker/api"
	"errors"

	"github.com/golang-jwt/jwt"
)

var UnvalidTokenError error = errors.New("The token is not valid")

func ValidateToken(token string, claims *api.Claims) (*jwt.Token, error) {

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if !tkn.Valid {
		return nil, UnvalidTokenError
	}
	return tkn, err
}
