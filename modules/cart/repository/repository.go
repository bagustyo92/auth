package repository

import (
	"github.com/bagustyo92/auth/modules/cart/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type cartRepo struct {
	gdb *gorm.DB
}

type CartRepoInterface interface {
	InsertCart(chart models.Cart) (*models.Cart, error)
	InsertProductCart(productCart models.ProductCart) error
	UpdateProductCart(productID uuid.UUID, dataUpdate map[string]interface{}) error
	GetCart(id uuid.UUID, filter models.ProductCartFilter) (*models.Cart, error)
	DeleteProductCart(cartID uuid.UUID, productCode string) error
}

func NewCartRepo(gdb *gorm.DB) CartRepoInterface {
	return &cartRepo{gdb}
}
