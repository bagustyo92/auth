package service

import (
	"errors"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/jinzhu/gorm"
)

func (as *authService) Create(auth models.Auth) (*models.Auth, error) {
	user, err := as.ar.GetByUsername(auth.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if auth.Password == "" {
				auth.Password = generateRandomPassword(4)
			}
			return as.ar.Create(auth)
		}

		return nil, err
	}

	return user, nil
}
