package main

import (
	jquants "Go-AutoTrade/j-quants"
	"Go-AutoTrade/utils"
	"fmt"
)

func main() {
	utils.InitLogger()
	JQClient, err := jquants.New()
	fmt.Println(JQClient, err)
}
