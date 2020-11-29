package repository

import (
	"fmt"

	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (ur *userRepo) InsertUser(user *models.User) error {
	return ur.gdb.Transaction(func(tx *gorm.DB) error {
		return ur.gdb.Create(user).Error
	})
}

func (ur *userRepo) GetUser(user *models.User) error {
	finalQuery := ur.gdb
	if user.UserName != "" {
		finalQuery = finalQuery.Where("user_name = ?", user.UserName)
	}
	finalQuery = finalQuery.Find(&user)
	return finalQuery.Error
}

func (ur *userRepo) GetUsers(query models.Query) (interface{}, error) {
	var users []models.User

	var finalQuery *gorm.DB
	fil := make(map[string]interface{})

	if query.UserID != nil {
		fil["user_id"] = query.UserID
	}

	finalQuery = ur.gdb.Where(fil)

	if query.Search != nil {
		finalQuery = finalQuery.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *query.Search))
	}

	pagination := pagination.Paging(&pagination.Param{
		DB:      finalQuery.Find(&users),
		Limit:   query.PageLimit,
		Page:    query.PageNumber,
		OrderBy: []string{"name asc"},
		ShowSQL: false,
	}, &users)

	return pagination, nil
}

func (ur *userRepo) UpdateUser(user *models.User) error {
	return ur.gdb.Transaction(func(tx *gorm.DB) error {
		if err := ur.gdb.Model(&user).Update(&user).Error; err != nil {
			return err
		}

		return ur.gdb.Find(&user).Error
	})
}

func (ur *userRepo) DeleteUser(id uuid.UUID) error {
	user := &models.User{}
	user.ID = id
	return ur.gdb.Delete(&user).Error
}
