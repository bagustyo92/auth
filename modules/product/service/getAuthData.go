package service

import "github.com/bagustyo92/auth/modules/auth/models"

func (cs *productService) GetAuthData(jwtClaim models.JWTClaims) (*models.Auth, error) {
	return cs.authService.GetByToken(jwtClaim)
}
