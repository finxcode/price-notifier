package entity

import "time"

type Quote struct {
	CoinId           int       `json:"coin_Id"`
	Symbol           string    `json:"symbol"`
	PercentChange24H float64   `json:"percentChange24H"`
	PercentChange7D  float64   `json:"percentChange7D"`
	LastUpdated      time.Time `json:"last_updated"`
}
