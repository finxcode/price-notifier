package out

import (
	"github.com/jmoiron/sqlx"
	"log"
	dataEntity "price-notifier/adapter/out/entity"
	"price-notifier/adapter/out/mapper"
	"price-notifier/application/domain/entity"
)

type CryptoAssetQuoteAdapter struct {
	db *sqlx.DB
}

func NewCryptoAssetQuoteAdapter(db *sqlx.DB) *CryptoAssetQuoteAdapter {
	return &CryptoAssetQuoteAdapter{
		db: db,
	}
}

func (c *CryptoAssetQuoteAdapter) GetQuoteByMarketCap(rank int, tradingDay string) []*entity.Quote {
	var quotes []dataEntity.Quote
	var quotesDomain []*entity.Quote
	query := "SELECT q.coin_id, c.symbol, q.percentChange24H, q.percentChange7D, q.last_updated " +
		"FROM quotes q LEFT JOIN coins c ON q.coin_id = c.coin_id " +
		"WHERE DATE_FORMAT(q.last_updated, '%Y-%m-%d') = $1 ORDER BY q.market_cap DESC LIMIT $2;"
	err := c.db.Select(&quotes, query, tradingDay, rank)
	if err != nil {
		log.Printf("query quotes db error: %s", err.Error())
	}
	if len(quotes) == 0 {
		log.Printf("no quote selected from db")
		return nil
	}

	for _, quote := range quotes {
		quotesDomain = append(quotesDomain, mapper.QuoteDataToDomain(&quote, tradingDay))
	}
	return quotesDomain
}

func (c *CryptoAssetQuoteAdapter) GetQuoteBySymbol(symbol string, tradingDay string) *entity.Quote {
	var quote dataEntity.Quote
	query := "SELECT q.coin_id, c.symbol, q.percentChange24H, q.percentChange7D, q.last_updated " +
		"FROM quotes q LEFT JOIN coins c ON q.coin_id = c.coin_id " +
		"WHERE DATE_FORMAT(q.last_updated, '%Y-%m-%d') = $1 AND c.symbol = $2;"
	err := c.db.Get(&quote, query, tradingDay, symbol)
	if err != nil {
		log.Printf("query quote of coin %s, db error: %s", symbol, err.Error())
		return nil
	}
	if quote.CoinId == 0 {
		log.Printf("no quote of coin %s found", symbol)
		return nil
	}
	return mapper.QuoteDataToDomain(&quote, tradingDay)
}
