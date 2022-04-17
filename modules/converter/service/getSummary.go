package service

import (
	"sort"
	"strconv"

	"github.com/bagustyo92/auth/modules/converter/models"
	reqModel "github.com/bagustyo92/auth/modules/request/models"
)

func converStringToInt64(a string) int32 {
	val, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}

	return int32(val)
}

func findMinimum(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

func findMax(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func findMedian(data []float64) float64 {
	sort.Float64s(data)

	medianIndex := len(data) / 2

	if len(data)%2 != 0 {
		return data[medianIndex]
	}

	return (data[medianIndex-1] + data[medianIndex]) / 2
}

func (cs *converterService) GetSummaryPriceList() (*models.SummaryPriceList, error) {
	var (
		sumarizeCity = map[string][]reqModel.Price{}

		summaryData models.SummaryPriceList
	)

	priceList, err := cs.GetPriceListIncludingPriceInUSD()
	if err != nil {
		return nil, err
	}

	// run through raw data
	for _, v := range priceList {
		if v.Province != "" {
			sumarizeCity[v.Province] = append(sumarizeCity[v.Province], v)
		}
	}

	// build sumarize data based on province resp
	for provinceName, datas := range sumarizeCity {
		var (
			min     int32
			max     int32
			median  float64
			average float32
			price   int32
			total   int32
			prices  []float64
		)

		for _, v := range datas {
			if v.Price != "" {
				price = converStringToInt64(v.Price)

				total += price
				prices = append(prices, float64(price))
			}
		}

		average = float32(total) / float32(len(datas))
		median = findMedian(prices)

		min = int32(prices[0])
		max = int32(prices[len(prices)-1])

		summaryData.SummaryBasedOnCity = append(summaryData.SummaryBasedOnCity, models.SummarizeCity{
			Province:     provinceName,
			MinPrice:     min,
			MaxPrice:     max,
			MedianPrice:  median,
			AveragePrice: average,
		})
	}

	return &summaryData, nil
}
