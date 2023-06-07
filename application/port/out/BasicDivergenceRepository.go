package out

import "price-notifier/application/domain/entity"

type BasicDivergenceRepository interface {
	Insert(divergence *entity.BasicDivergence) error
	//GetCoinsByDivergenceScore(rank int, tradingDay string)
}
