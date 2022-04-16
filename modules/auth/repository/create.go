package repository

import (
	"github.com/bagustyo92/auth/modules/auth/models"
)

func (ar *authRepo) Create(auth models.Auth) (*models.Auth, error) {
	if err := ar.gdb.Create(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
