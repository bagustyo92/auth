package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/bagustyo92/auth/modules/user/models"
	uuid "github.com/satori/go.uuid"
)

type Auth struct {
	models.Base

	UserID      uuid.UUID `json:"userID,omitempty"`
	AccessToken string    `json:"lastAccessToken,omitempty" `
	LastLogin   time.Time `json:"lastLogin"`
	IP          string    `json:"ip"`
}

type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Claim jwt.StandardClaims
}
