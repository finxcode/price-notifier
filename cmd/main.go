package main

import (
	cron "github.com/robfig/cron/v3"
	"log"
	"os"
	"price-notifier/adapter/in"
	"price-notifier/adapter/out"
	"price-notifier/application/config"
	"price-notifier/application/service"
	"strconv"
	"time"
)

func main() {
	rank, err := strconv.Atoi(os.Getenv("RANK"))
	if err != nil {
		log.Fatalln("rank should be an integer")
	}
	baselineAssetSymbol := os.Getenv("BASELINE_ASSET_SYMBOL")
	fetchTime := time.Now()
	c := config.NewConfig(rank, baselineAssetSymbol, fetchTime)

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	dbConfig := NewConfig(user, password, host, port, database)

	db, err := InitDB(dbConfig)
	if err != nil {
		log.Fatalf("database connection error: %s", err.Error())
	}

	cryptoAssetQuoteAdapter := out.NewCryptoAssetQuoteAdapter(db)
	divergenceAdapter := out.NewDivergenceAdapter(db)
	basicDivergenceService := service.NewBasicDivergenceService(c, cryptoAssetQuoteAdapter, divergenceAdapter)
	basicDivergenceAdapter := in.NewBasicDivergenceAdapter(basicDivergenceService)

	cronJob := cron.New()
	entryID, err := cronJob.AddFunc("45 14 * * *", func() {
		basicDivergenceAdapter.StoreBasicDivergences()
	})
	if err != nil {
		log.Printf("Divergence job executed on %v with entryId: %v and with error: %s",
			time.Now(), entryID, err.Error())
	} else {
		log.Printf("Divergence job executed on %v with entryId: %v",
			time.Now(), entryID)
	}

	cronJob.Start()
	select {}
}
