package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid)
}

type User struct {
	Base

	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
}

type Query struct {
	PageLimit  int        `json:"pageLimit"`
	PageNumber int        `json:"pageNumber"`
	Search     *string    `json:"search"`
	UserID     *uuid.UUID `json:"userID"`
}
