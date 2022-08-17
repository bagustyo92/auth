package service

import (
	"github.com/bagustyo92/auth/modules/cart/models"
	"github.com/bagustyo92/auth/modules/cart/repository"
	uuid "github.com/satori/go.uuid"
)

type cartService struct {
	ur repository.CartRepoInterface
}

type CartInterface interface {
	AddProductToCart(cartReq models.Cart) (*models.Cart, error)
	GetCart(cartID uuid.UUID, filter models.ProductCartFilter) (*models.Cart, error)
	DeleteProductFromCart(cartID uuid.UUID, productCode string) error
	MoneyTest(amountOfMoney int) map[string]int
	StringTest(text1, text2 string) bool
}

func NewCartService(cartRepo repository.CartRepoInterface) CartInterface {
	return &cartService{cartRepo}
}
