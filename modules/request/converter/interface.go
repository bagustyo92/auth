package converter

import (
	"net/http"
)

type Interface interface {
	GetCurrUSD(header *http.Header, usd **float64) error
}
