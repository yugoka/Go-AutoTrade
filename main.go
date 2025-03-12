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

	res, _ := jqClient.GetStatements(jquants.GetStatementsParams{
		Code: "9104",
	})

	fmt.Println(jquants.GenerateStatementsReport(res))
}
