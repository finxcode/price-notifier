package in

import (
	"log"
	"price-notifier/application/domain/entity"
	"price-notifier/application/port/in"
)

type BasicDivergenceAdapter struct {
	basicDivergencePort in.BasicDivergencePort
}

func NewBasicDivergenceAdapter(basicDivergencePort in.BasicDivergencePort) *BasicDivergenceAdapter {
	return &BasicDivergenceAdapter{
		basicDivergencePort: basicDivergencePort,
	}
}

func (p *BasicDivergenceAdapter) CalcBasicDivergence() []*entity.BasicDivergence {
	divergences := p.basicDivergencePort.CalcBasicDivergence()
	log.Printf("%d coin basic divergence has been calulated", len(divergences))
	return divergences
}

func (p *BasicDivergenceAdapter) StoreBasicDivergences() {
	divergences := p.CalcBasicDivergence()
	numStored := p.basicDivergencePort.SaveBasicDivergence(divergences)
	log.Printf("%d coin basic divergence has been stored", numStored)
}
