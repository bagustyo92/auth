package efishery

import (
	"net/http"

	"github.com/bagustyo92/auth/modules/request/models"
)

type Interface interface {
	GetPriceLists(header *http.Header, priceList **[]models.Price) error
}
