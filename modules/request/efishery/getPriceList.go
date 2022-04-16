package efishery

import (
	"net/http"

	"github.com/bagustyo92/auth/config"
	"github.com/bagustyo92/auth/modules/request"
	"github.com/bagustyo92/auth/modules/request/models"
)

func (r *req) GetPriceLists(header *http.Header, priceList **[]models.Price) error {
	reqParam := request.RequestParams{
		Url:     config.EfisheryAPIHost + "/v1/storages/5e1edf521073e315924ceab4/list",
		Method:  http.MethodGet,
		Header:  header,
		Payload: nil,
	}

	if err := request.Request(&reqParam, &priceList); err != nil {
		return err
	}

	return nil
}
