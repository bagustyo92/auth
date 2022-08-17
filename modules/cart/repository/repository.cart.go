package repository

import (
	"fmt"
	"strings"

	"github.com/bagustyo92/auth/modules/cart/models"
	uuid "github.com/satori/go.uuid"
)

func (ur *cartRepo) InsertCart(cart models.Cart) (*models.Cart, error) {
	if err := ur.gdb.Create(cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (ur *cartRepo) GetCart(id uuid.UUID, filter models.ProductCartFilter) (*models.Cart, error) {
	cart := models.Cart{}
	products := []models.ProductCart{}

	if (models.ProductCartFilter{}) == filter {
		if err := ur.gdb.Where("id = ?", id).Preload("Products").First(&cart).Error; err != nil {
			return nil, err
		}
	} else {
		if err := ur.gdb.Where("id = ?", id).First(&cart).Error; err != nil {
			return nil, err
		}

		queryProducts := ur.gdb.Where("cart_id = ?", id)

		if filter.ProductName != "" {
			queryProducts = queryProducts.Where(fmt.Sprintf("LOWER(product_name) LIKE '%%%s%%'", strings.ToLower(filter.ProductName)))
		}

		if filter.Quantity != "" {
			queryProducts = queryProducts.Where("quantity = ?", filter.Quantity)
		}

		if err := queryProducts.Find(&products).Error; err != nil {
			return nil, err
		}
		cart.Products = products
	}

	return &cart, nil
}

func (ur *cartRepo) InsertProductCart(productCart models.ProductCart) error {
	return ur.gdb.Create(productCart).Error
}

func (ur *cartRepo) UpdateProductCart(productID uuid.UUID, dataUpdate map[string]interface{}) error {
	return ur.gdb.Model(&models.ProductCart{}).
		Where("id = ?", productID).
		Updates(dataUpdate).Error
}

func (ur *cartRepo) DeleteProductCart(cartID uuid.UUID, productCode string) error {
	fmt.Println(cartID, productCode)
	return ur.gdb.
		Where("cart_id = ?", cartID).
		Where("product_code = ?", productCode).
		Delete(&models.ProductCart{}).
		Error
}
