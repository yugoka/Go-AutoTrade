package utils

import (
	"Go-AutoTrade/config"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// ログ保持件数
const maxLogFiles = 10

func InitLogger() {
	logDir := config.GlobalConfig.LogOutputPath
	if logDir == "" {
		log.Println("LOG_OUTPUT_PATH is not set, using default stdout")
		return
	}

	// 実行ごとに異なるログファイル名を生成 (app_YYYY-MM-DD_HH-MM-SS.log)
	logFileName := filepath.Join(logDir, "app_"+time.Now().Format("2006-01-02_15-04-05")+".log")

	// ログファイルを開く
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// ログ出力先をファイルに設定
	log.SetOutput(logFile)
	log.Println("Logger initialized: writing to", logFileName)

	// 古いログを削除
	cleanupOldLogs(logDir)
}

func cleanupOldLogs(logDir string) {
	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Printf("Failed to read log directory: %v", err)
		return
	}

	var logFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "app_") && strings.HasSuffix(file.Name(), ".log") {
			logFiles = append(logFiles, file.Name())
		}
	}

	sort.Slice(logFiles, func(i, j int) bool {
		return logFiles[i] < logFiles[j] // 文字列比較で時系列順に並ぶ
	})

	// 上限以上のログがあれば古いものから削除
	if len(logFiles) > maxLogFiles {
		for _, oldLog := range logFiles[:len(logFiles)-maxLogFiles] {
			oldLogPath := filepath.Join(logDir, oldLog)
			err := os.Remove(oldLogPath)
			if err != nil {
				log.Printf("Failed to remove old log file %s: %v", oldLogPath, err)
			} else {
				log.Printf("Deleted old log file: %s", oldLogPath)
			}
		}
	}
}
