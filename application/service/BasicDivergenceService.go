package service

import (
	"log"
	"price-notifier/application/config"
	"price-notifier/application/domain/entity"
	"price-notifier/application/port/out"
)

type SevenDaysPriceDivergenceService struct {
	c                          *config.Config
	cryptoAssetQuoteRepository out.CryptoAssetQuoteRepository
	basicDivergenceRepository  out.BasicDivergenceRepository
}

func NewBasicDivergenceService(
	config *config.Config,
	cryptoAssetQuoteRepository out.CryptoAssetQuoteRepository,
	basicDivergenceRepository out.BasicDivergenceRepository) *SevenDaysPriceDivergenceService {
	return &SevenDaysPriceDivergenceService{
		c:                          config,
		cryptoAssetQuoteRepository: cryptoAssetQuoteRepository,
		basicDivergenceRepository:  basicDivergenceRepository,
	}
}

func (s *SevenDaysPriceDivergenceService) CalcBasicDivergence() []*entity.BasicDivergence {
	var divergences []*entity.BasicDivergence
	quotes := s.cryptoAssetQuoteRepository.GetQuoteByMarketCap(s.c.Rank, TimeToTradingDay(s.c.FetchTime))
	baselineAssetQuote := s.cryptoAssetQuoteRepository.GetQuoteBySymbol(s.c.BaselineAssetSymbol, TimeToTradingDay(s.c.FetchTime))
	for _, quote := range quotes {
		basicDivergence, err := entity.NewDivergence(quote, baselineAssetQuote)
		if err != nil {
			log.Printf("an error occurred when calc divergence on coin %s: %s", quote.Symbol, err.Error())
		}
		divergences = append(divergences, basicDivergence)
		if err := s.basicDivergenceRepository.Insert(basicDivergence); err != nil {
			log.Printf("insert divergence for coin %s falied with error: %s", quote.Symbol, err.Error())
		}
	}
	return divergences
}

func (s *SevenDaysPriceDivergenceService) SaveBasicDivergence(divergences []*entity.BasicDivergence) int {
	suc := 0
	for _, divergence := range divergences {
		if err := s.basicDivergenceRepository.Insert(divergence); err != nil {
			log.Printf("insert divergence for coin %d falied with error: %s", divergence.CoinId, err.Error())
		} else {
			suc++
		}
	}
	return suc
}
