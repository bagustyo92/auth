package service

import (
	"time"

	"github.com/bagustyo92/auth/modules/cart/models"
	uuid "github.com/satori/go.uuid"
)

func validateQty(qty int) int {
	if qty == 0 {
		return 1
	}
	return qty
}

func (us *cartService) AddProductToCart(cartReq models.Cart) (*models.Cart, error) {
	if cartReq.ID == uuid.Nil {
		tn := time.Now()
		// Insert New
		cartReq.ID = uuid.NewV4()
		cartReq.CreatedAt = &tn
		cartReq.UpdatedAt = &tn

		for i, v := range cartReq.Products {
			cartReq.Products[i].Quantity = validateQty(v.Quantity)
		}
		return us.ur.InsertCart(cartReq)
	} else {
		// Get Cart
		cart, err := us.ur.GetCart(cartReq.ID, models.ProductCartFilter{})
		if err != nil {
			return nil, err
		}

		// is product exist
		existProduct := make(map[string]int)
		existProductID := make(map[string]uuid.UUID)
		for _, product := range cart.Products {
			existProduct[product.ProductCode] = product.Quantity
			existProductID[product.ProductCode] = product.ID
		}

		for _, productReq := range cartReq.Products {
			_, isExist := existProduct[productReq.ProductCode]
			if isExist {
				// Update Qty
				if err := us.ur.UpdateProductCart(existProductID[productReq.ProductCode], map[string]interface{}{
					"quantity": existProduct[productReq.ProductCode] + validateQty(productReq.Quantity),
				}); err != nil {
					return nil, err
				}
			} else {
				// Insert New One
				tn := time.Now()
				productReq.ID = uuid.NewV4()
				productReq.CartID = cart.ID
				productReq.CreatedAt = &tn
				productReq.UpdatedAt = &tn
				productReq.Quantity = validateQty(productReq.Quantity)
				if err := us.ur.InsertProductCart(productReq); err != nil {
					return nil, err
				}
			}
		}
	}

	return nil, nil
}

func (us *cartService) GetCart(cartID uuid.UUID, filter models.ProductCartFilter) (*models.Cart, error) {
	return us.ur.GetCart(cartID, filter)
}

func (us *cartService) DeleteProductFromCart(cartID uuid.UUID, productCode string) error {
	return us.ur.DeleteProductCart(cartID, productCode)
}

func (us *cartService) MoneyTest(amountOfMoney int) map[string]int {
	moneys := gatheringMoneyOnRupiah(amountOfMoney)
	return makeMoneyResponse(moneys)
}

func (us *cartService) StringTest(text1, text2 string) bool {
	if stringEditAssessment(text1, text2) < 2 {
		return true
	} else {
		return false
	}
}

func makeMoneyResponse(moneyGathered map[int]int) map[string]int {
	indonesiaMoney := map[int]string{
		100000: "Rp 100.000,-",
		50000:  "Rp 50.000,-",
		20000:  "Rp 20.000,-",
		10000:  "Rp 10.000,-",
		5000:   "Rp 5.000,-",
		2000:   "Rp 2.000,-",
		1000:   "Rp 1.000,-",
		500:    "Rp 500,-",
		200:    "Rp 200,-",
		100:    "Rp 100,-",
	}

	resp := make(map[string]int)

	for i, v := range moneyGathered {
		resp[indonesiaMoney[i]] = v
	}

	return resp
}

func gatheringMoneyOnRupiah(moneyInput int) map[int]int {
	if moneyInput == 0 {
		return nil
	}

	moneyList := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	moneyGathered := make(map[int]int)
	i := 0

	for {
		if moneyInput >= moneyList[i] {
			moneyInput -= moneyList[i]
			moneyGathered[moneyList[i]] += 1
		} else {
			if len(moneyList)-1 == i {
				moneyGathered[moneyList[i]] += 1
				break
			}
			i += 1
		}

		if moneyInput == 0 {
			break
		}
	}

	return moneyGathered
}

func stringEditAssessment(text1, text2 string) int {
	var (
		totalStep int
	)
	wordCounter := make(map[rune]int)

	for _, r := range text1 {
		wordCounter[r] += 1
	}

	for _, r := range text2 {
		wordCounter[r] += 1
	}

	for _, v := range wordCounter {
		if v%2 != 0 {
			totalStep += 1
		}
	}

	return totalStep
}
