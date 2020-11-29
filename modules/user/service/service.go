package service

import (
	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/bagustyo92/auth/modules/user/repository"
	uuid "github.com/satori/go.uuid"
)

type userService struct {
	ur repository.UserRepoInterface
}

type UserServiceInterface interface {
	CreateUser(user *models.User) error
	GetUser(user *models.User) error
	GetUsers(query models.Query) (interface{}, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

func NewUserService(userRepo repository.UserRepoInterface) UserServiceInterface {
	return &userService{userRepo}
}
