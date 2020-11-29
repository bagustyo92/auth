package service

import (
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/modules/auth/repository"
	userModels "github.com/bagustyo92/auth/modules/user/models"
	userRepo "github.com/bagustyo92/auth/modules/user/repository"
	uuid "github.com/satori/go.uuid"
)

type authService struct {
	ar repository.AuthRepoInterface
	ur userRepo.UserRepoInterface
}

type AuthServiceInterface interface {
	Login(authString string) (*userModels.User, error)
	CreateJwtToken(user *userModels.User) (string, error)

	CreateAuth(auth *models.Auth) error
	UpdateAuth(auth *models.Auth) error
	GetAuth(userID uuid.UUID) (*models.Auth, error)
}

func NewAuthService(authRepo repository.AuthRepoInterface, userRepo userRepo.UserRepoInterface) AuthServiceInterface {
	return &authService{authRepo, userRepo}
}
