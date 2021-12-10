package main

import (
	"github.com/docker/docker/api/types"
	"github.com/golang-jwt/jwt"
)

const (
	JWTKey    = "docker-watchdog.email.cache"
	JWTIssuer = "docker-watchdog"
)

type JWTPayload struct {
	jwt.StandardClaims
	Containers []JWTContainersData `json:"containers"`
}

type JWTContainersData struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

func encodeToJWT(containers []types.Container) (*string, error) {
	claims := JWTPayload{}

	for _, c := range containers {
		claims.Containers = append(claims.Containers, JWTContainersData{
			ID:    c.ID[:10],
			State: c.State,
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func decodeJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
