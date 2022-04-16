package repository

import (
	"github.com/bagustyo92/auth/modules/auth/models"
)

type AuthRepoInterface interface {
	Create(auth models.Auth) (*models.Auth, error)
	GetByUsername(username string) (*models.Auth, error)
	GetByPhone(phoneNumber string) (*models.Auth, error)
}
