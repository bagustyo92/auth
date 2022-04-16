package service

import (
	"github.com/bagustyo92/auth/modules/auth/models"
)

type AuthServiceInterface interface {
	Login(userLogin models.Auth) (string, error)
	Create(auth models.Auth) (*models.Auth, error)
	GetByToken(tokenClaim models.JWTClaims) (*models.Auth, error)
}
