package app

import (
	"fmt"
	"task_2/internal/services/http"
	"task_2/internal/services/logger"
	"task_2/internal/services/reader"
	"task_2/internal/services/sum"
)

type App struct {
	reader        reader.Reader
	sumCalculator sum.Calculator
	httpChecker   http.Checker
	logger        logger.Logger
	url           string
}

func NewApp(r reader.Reader, s sum.Calculator, h http.Checker, l logger.Logger, url string) *App {
	return &App{
		reader:        r,
		sumCalculator: s,
		httpChecker:   h,
		logger:        l,
		url:           url,
	}
}

func (a *App) Run() {
	data, err := a.reader.ReadData()
	if err != nil {
		a.logger.LogError("Failed to read data", err)
		return
	}

	totalSum := a.sumCalculator.SumNumbers(data)
	a.logger.LogInfo(fmt.Sprintf("Total sum: %d", totalSum))
	fmt.Printf("Total sum: %d\n", totalSum)
	status, err := a.httpChecker.CheckURL(a.url)
	if err != nil {
		a.logger.LogError("Failed to check URL", err)
		return
	}

	a.logger.LogInfo(fmt.Sprintf("URL status: %d", status))
	fmt.Printf("URL status: %d\n", status)
}
