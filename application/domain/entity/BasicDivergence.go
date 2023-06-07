package entity

import "errors"

var NoAssetErr = errors.New("no asset or baseline passed in")

type BasicDivergence struct {
	CoinId          int
	Symbol          string
	BaselineCoinId  int
	Divergence24H   float64
	Divergence7D    float64
	DivergenceTotal float64
	TradingDay      string
}

func NewDivergence(asset, baseline *Quote) (*BasicDivergence, error) {
	if asset == nil || baseline == nil {
		return nil, NoAssetErr
	}

	divergence24H := asset.PercentChange24H - baseline.PercentChange24H
	divergence7D := asset.PercentChange7D - baseline.PercentChange7D
	divergenceTotal := divergence24H + divergence7D
	return &BasicDivergence{
		CoinId:          asset.CoinId,
		Symbol:          asset.Symbol,
		BaselineCoinId:  baseline.CoinId,
		Divergence24H:   divergence24H,
		Divergence7D:    divergence7D,
		DivergenceTotal: divergenceTotal,
		TradingDay:      baseline.TradingDay,
	}, nil
}
