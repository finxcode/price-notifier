package service

import "time"

func TimeToTradingDay(t time.Time) string {
	return t.String()[0:10]
}
