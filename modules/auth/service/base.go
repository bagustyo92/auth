package service

import "github.com/bagustyo92/auth/modules/auth/repository"

type authService struct {
	ar repository.AuthRepoInterface
}

func NewAuthService(authRepo repository.AuthRepoInterface) AuthServiceInterface {
	return &authService{authRepo}
}
