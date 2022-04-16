package models

import "time"

type SummaryPriceList struct {
	SummaryBasedOnCity []SummarizeCity `json:"summary_based_on_city"`
	SummaryBasedOnDate []SummarizeDate `json:"summary_based_on_date"`
}

type SummarizeDate struct {
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	MinPrice     int32     `json:"min_price"`
	MaxPrice     int32     `json:"max_price"`
	MedianPrice  int32     `json:"median_price"`
	AveragePrice int32     `json:"average_price"`
}

type SummarizeCity struct {
	Province     string  `json:"area_provinsi"`
	MinPrice     int32   `json:"min_price"`
	MaxPrice     int32   `json:"max_price"`
	MedianPrice  float64 `json:"median_price"`
	AveragePrice float32 `json:"average_price"`
}
