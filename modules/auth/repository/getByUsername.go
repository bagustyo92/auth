package repository

import (
	"github.com/bagustyo92/auth/modules/auth/models"
)

func (ar *authRepo) GetByUsername(username string) (*models.Auth, error) {
	auth := models.Auth{}

	if err := ar.gdb.Where("username = ?", username).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
