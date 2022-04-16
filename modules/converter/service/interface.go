package service

import (
	authModel "github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/modules/converter/models"
	reqModel "github.com/bagustyo92/auth/modules/request/models"
)

type ConverterServiceInterface interface {
	GetPriceListIncludingPriceInUSD() ([]reqModel.Price, error)
	GetSummaryPriceList() (*models.SummaryPriceList, error)
	GetAuthData(jwtClaim authModel.JWTClaims) (*authModel.Auth, error)
}
