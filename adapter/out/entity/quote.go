package entity

import "time"

type Quote struct {
	CoinId           int       `json:"coin_id" db:"coin_id"`
	Symbol           string    `json:"symbol" db:"symbol"`
	PercentChange24H float64   `json:"percentChange24H" db:"percentChange24H"`
	PercentChange7D  float64   `json:"percentChange7D" db:"percentChange7D"`
	LastUpdated      time.Time `json:"last_updated" db:"last_updated"`
}
