package service

import (
	"net/http"
	"strconv"

	"github.com/bagustyo92/auth/modules/request/models"
)

func (cs *converterService) GetPriceListIncludingPriceInUSD() ([]models.Price, error) {
	var (
		aDollar   *float64
		header    = http.Header{}
		priceList *[]models.Price
	)

	header.Set("Content-Type", "application/json")

	// TODO: Store to local mem cache

	// get current usd to idr
	if err := cs.currencyConverter.GetCurrUSD(&header, &aDollar); err != nil {
		return nil, err
	}

	// fetch data from efishery
	if err := cs.efishery.GetPriceLists(&header, &priceList); err != nil {
		return nil, err
	}

	for i, price := range *priceList {
		if price.Price != "" {
			realPrice, err := strconv.Atoi(price.Price)
			if err != nil {
				continue
			}

			(*priceList)[i].PriceUSD = float64(realPrice) / *aDollar
		}
	}

	return *priceList, nil
}
