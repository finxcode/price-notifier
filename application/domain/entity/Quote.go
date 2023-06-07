package entity

type Quote struct {
	CoinId           int
	Symbol           string
	PercentChange24H float64
	PercentChange7D  float64
	TradingDay       string
}

func NewQuote(coinId int, symbol string, percentChange24H, percentChange7D float64, tradingDay string) *Quote {
	return &Quote{
		CoinId:           coinId,
		Symbol:           symbol,
		PercentChange7D:  percentChange7D,
		PercentChange24H: percentChange24H,
		TradingDay:       tradingDay,
	}
}
