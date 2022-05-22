package jwtToken

import (
	"BugTracker/api"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey []byte = []byte("davidLavariete")

func GenerateToken(username string, isRefreshToken bool) (string, error) {
	tknType := "validation"
	if isRefreshToken {
		tknType = "refresh"
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &api.Claims{
		Username: username,
		Type:     tknType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}).SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
