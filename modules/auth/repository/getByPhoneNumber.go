package repository

import "github.com/bagustyo92/auth/modules/auth/models"

func (ar *authRepo) GetByPhone(phoneNumber string) (*models.Auth, error) {
	auth := models.Auth{}

	if err := ar.gdb.First(&auth).Where("phone = ?", phoneNumber).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
