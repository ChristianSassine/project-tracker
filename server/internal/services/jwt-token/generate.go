package jwtToken

import (
	"os"
	"strconv"
	"time"

	"github.com/ChristianSassine/projectManager/internal/api"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey []byte = []byte(os.Getenv("JWT_KEY"))

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
