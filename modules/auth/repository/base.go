package repository

import (
	"github.com/jinzhu/gorm"
)

type authRepo struct {
	gdb *gorm.DB
}

func NewRepository(gdb *gorm.DB) AuthRepoInterface {
	return &authRepo{gdb}
}
