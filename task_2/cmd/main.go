package main

import (
	"log"
	"task_2/internal/app"
	"task_2/internal/config"
	"task_2/internal/services/http"
	"task_2/internal/services/logger"
	"task_2/internal/services/reader"
	"task_2/internal/services/sum"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := app.NewApp(
		reader.NewFileReader(cfg.InputFile),
		sum.NewSimpleSumCalculator(),
		http.NewSimpleHTTPChecker(),
		logger.NewFileLogger(cfg.LogFile),
		cfg.URL,
	)

	app.Run()
}
