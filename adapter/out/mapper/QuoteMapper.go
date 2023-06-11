package mapper

import (
	dataEntity "price-notifier/adapter/out/entity"
	"price-notifier/application/domain/entity"
)

func QuoteDataToDomain(quoteData *dataEntity.Quote, tradingDay string) *entity.Quote {
	return entity.NewQuote(quoteData.CoinId, quoteData.Symbol, quoteData.PercentChange24H,
		quoteData.PercentChange7D, tradingDay)
}
