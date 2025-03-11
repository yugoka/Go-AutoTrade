package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GlobalConfigList struct {
	LogOutputPath string
}
var GlobalConfig GlobalConfigList

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}
}

func init() {
	loadEnv()
	
	GlobalConfig = GlobalConfigList {
		LogOutputPath: os.Getenv("LOG_OUTPUT_PATH"),
	}
}