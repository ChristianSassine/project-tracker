package jwtToken

import (
	"BugTracker/api"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey []byte = []byte("davidLavariete")

func GenerateToken(username string, id int, expiryTime time.Duration, isRefreshToken bool) (string, error) {
	tknType := "validation"
	if isRefreshToken {
		tknType = "refresh"
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &api.Claims{
		Username: username,
		Type:     tknType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiryTime).Unix(),
			Subject:   strconv.Itoa(id),
		},
	}).SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
