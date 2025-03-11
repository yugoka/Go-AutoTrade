package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GlobalConfigList struct {
	LogOutputPath string
	JQuantsMailAddress string
	JQuantsPassword string
}
var GlobalConfig GlobalConfigList

func loadEnv() {
	if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found or could not be loaded.")
	}
}

func init() {
	loadEnv()

	GlobalConfig = GlobalConfigList {
		LogOutputPath: os.Getenv("LOG_OUTPUT_PATH"),
		JQuantsMailAddress: os.Getenv("J_QUANTS_MAIL_ADDRESS"),
		JQuantsPassword: os.Getenv("J_QUANTS_PASSWORD"),
	}
}