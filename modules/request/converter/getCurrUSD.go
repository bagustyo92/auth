package converter

import (
	"net/http"

	"github.com/bagustyo92/auth/config"
	"github.com/bagustyo92/auth/modules/request"
)

func (r *req) GetCurrUSD(header *http.Header, usd **float64) error {
	var (
		resp map[string]interface{}
	)
	reqParam := request.RequestParams{
		Url:     config.CurrencyConverterAPIHost + "/api/v7/convert?q=USD_IDR&&compact=ultra&apiKey=" + config.CurrencyConverterAPIKey,
		Method:  http.MethodGet,
		Header:  header,
		Payload: nil,
	}

	if err := request.Request(&reqParam, &resp); err != nil {
		return err
	}

	aDollar := resp["USD_IDR"].(float64)
	*usd = &aDollar

	return nil
}
