package service

import (
	"net/http"
	"strconv"

	"github.com/bagustyo92/auth/config"
	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/request/models"
)

const (
	usdToIdrCacheKey = "get.usd.to.idr.converter"
)

func (cs *converterService) GetPriceListIncludingPriceInUSD() ([]models.Price, error) {
	var (
		aDollar   *float64
		header    = http.Header{}
		priceList *[]models.Price
	)

	header.Set("Content-Type", "application/json")

	// Check from local memcache
	dollar, err := cs.cacher.Get(usdToIdrCacheKey)
	if err != nil {
		// get current usd to idr
		if err := cs.currencyConverter.GetCurrUSD(&header, &aDollar); err != nil {
			return nil, err
		}

		cs.cacher.Set(usdToIdrCacheKey, aDollar, int64(config.CacheDefaultTTL))
	} else {
		aDollar = dollar.(*float64)
		logger.MakeLogEntry(nil).Info("Data currency provide by memlocal cache. Value is ", *aDollar)
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
