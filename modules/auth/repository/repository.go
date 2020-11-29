package repository

import (
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type authRepo struct {
	gdb *gorm.DB
}

type AuthRepoInterface interface {
	CreateAuth(auth *models.Auth) error
	UpdateAuth(auth *models.Auth) error
	GetAuth(userID uuid.UUID) (*models.Auth, error)
}
