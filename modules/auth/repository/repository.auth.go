package repository

import (
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func NewAuthRepo(gdb *gorm.DB) AuthRepoInterface {
	return &authRepo{gdb}
}

func (ar *authRepo) CreateAuth(auth *models.Auth) error {
	return ar.gdb.Create(&auth).Error
}

func (ar *authRepo) UpdateAuth(auth *models.Auth) error {
	return ar.gdb.Model(&auth).Update(&auth).Error
}

func (ar *authRepo) GetAuth(userID uuid.UUID) (*models.Auth, error) {
	auth := models.Auth{}
	auth.ID = userID

	if err := ar.gdb.Find(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
