package models

type Price struct {
	UUID       string  `json:"uuid"`
	Commodity  string  `json:"komoditas"`
	Province   string  `json:"area_provinsi"`
	City       string  `json:"area_kota"`
	Size       string  `json:"size"`
	Price      string  `json:"price"`
	PriceUSD   float64 `json:"price_usd"`
	TimeParsed string  `json:"tgl_parsed"`
	Timestamp  string  `json:"timestamp"`
}
