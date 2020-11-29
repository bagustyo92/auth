package service

import (
	"errors"

	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/bagustyo92/auth/utils"
	uuid "github.com/satori/go.uuid"
)

func (us *userService) CreateUser(user *models.User) error {
	if err := us.ur.GetUser(user); err == nil {
		return errors.New("username already exist, please used another username")
	}

	pwd, err := utils.SetPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	return us.ur.InsertUser(user)
}

func (us *userService) GetUser(user *models.User) error {
	return us.ur.GetUser(user)
}

func (us *userService) GetUsers(query models.Query) (interface{}, error) {
	return us.ur.GetUsers(query)
}

func (us *userService) UpdateUser(user *models.User) error {
	if user.Password != "" {
		pwd, err := utils.SetPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = pwd
	}

	return us.ur.UpdateUser(user)
}

func (us *userService) DeleteUser(id uuid.UUID) error {
	return us.ur.DeleteUser(id)
}
