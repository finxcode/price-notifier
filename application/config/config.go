package config

import "time"

type Config struct {
	Rank                int
	BaselineAssetSymbol string
	FetchTime           time.Time
}

func NewConfig(rank int, baselineAssetSymbol string, fetchTime time.Time) *Config {
	return &Config{
		Rank:                rank,
		BaselineAssetSymbol: baselineAssetSymbol,
		FetchTime:           fetchTime,
	}
}

func (c *Config) SetFetchTime(fetchTime time.Time) {
	c.FetchTime = fetchTime
}
