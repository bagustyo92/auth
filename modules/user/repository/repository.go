package repository

import (
	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type userRepo struct {
	gdb *gorm.DB
}

type UserRepoInterface interface {
	InsertUser(user *models.User) error
	GetUser(user *models.User) error
	GetUsers(query models.Query) (interface{}, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

func NewUserRepo(gdb *gorm.DB) UserRepoInterface {
	return &userRepo{gdb}
}
