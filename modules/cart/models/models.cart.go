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

type ProductCart struct {
	Base

	ProductName string    `json:"product_name"`
	ProductCode string    `json:"product_code"`
	Quantity    int       `json:"quantity"`
	CartID      uuid.UUID `json:"cart_id"`
}

type Cart struct {
	Base

	Products []ProductCart `json:"products" grom:"foreignKey:CartID"`
}

type ProductCartFilter struct {
	ProductName string `query:"product_name"`
	Quantity    string `query:"quantity"`
}
