package out

import "price-notifier/application/domain/entity"

type CryptoAssetQuoteRepository interface {
	GetQuoteByMarketCap(rank int, tradingDay string) []*entity.Quote
	GetQuoteBySymbol(symbol string, tradingDay string) *entity.Quote
}
