package config

import (
	"blogv2/model"

	"github.com/golang-jwt/jwt/v5"
)

type PayloadClaims struct {
	Payload model.UserPayload `json:"payload"`
	jwt.RegisteredClaims
}