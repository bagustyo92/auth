package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Base struct {
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Auth struct {
	Base

	ID       uint64 `json:"id" gorm:"primaryKey"`
	Phone    string `json:"phone" gorm:"index:idx_phone,unique" validate:"required"`
	Username string `json:"username" gorm:"index:idx_username,unique" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken string `json:"accessToken"`
}

type JWTClaims struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt int64     `json:"exp,omitempty"`
	Id        string    `json:"jti,omitempty"`
	IssuedAt  int64     `json:"iat,omitempty"`
	Issuer    string    `json:"iss,omitempty"`
	jwt.StandardClaims
}
