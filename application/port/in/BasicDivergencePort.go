package in

import "price-notifier/application/domain/entity"

type BasicDivergencePort interface {
	CalcBasicDivergence() []*entity.BasicDivergence
	SaveBasicDivergence([]*entity.BasicDivergence) int
}
