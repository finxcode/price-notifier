package out

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"price-notifier/application/domain/entity"
)

type DivergenceAdapter struct {
	db *sqlx.DB
}

func NewDivergenceAdapter(db *sqlx.DB) *DivergenceAdapter {
	return &DivergenceAdapter{
		db: db,
	}
}

func (d *DivergenceAdapter) Insert(divergence *entity.BasicDivergence) error {
	query := "INSERT " +
		"INTO basic_divergences " +
		"(coin_id, baseline_id,symbol,divergence24H,divergence7D,divergence_total,trading_day)" +
		"VALUES (?,?,?,?,?,?) "
	res := d.db.MustExec(query, divergence.CoinId, divergence.BaselineCoinId,
		divergence.Divergence24H, divergence.Divergence7D, divergence.DivergenceTotal, divergence.TradingDay)
	if _, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("db error:%s", err.Error())
	}
	return nil
}
