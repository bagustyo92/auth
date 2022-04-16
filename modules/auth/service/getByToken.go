package service

import "github.com/bagustyo92/auth/modules/auth/models"

func (as *authService) GetByToken(tokenClaim models.JWTClaims) (*models.Auth, error) {
	return as.ar.GetByUsername(tokenClaim.Name)
}
