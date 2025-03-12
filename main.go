package main

import (
	jquants "Go-AutoTrade/j-quants"
	"Go-AutoTrade/utils"
	"fmt"
	"log"
)

func main() {
	utils.InitLogger()
	jqClient, err := jquants.New()
	if err != nil {
		log.Fatalf("Failed to init JQuantsClient: %s", err)
	}

	res, _ := jqClient.GetDailyQuotes(jquants.GetDailyQuotesParams{
		Code: "4478",
		From: "20220101",
		To:   "20221230",
	})

	for _, data := range res {
		fmt.Println((data.High + data.Low) / 2)
	}
}
