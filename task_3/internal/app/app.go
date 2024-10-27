package app

import (
	"fmt"
	"log"
	"os"
	"sync"

	"task_2/internal/config"
	"task_2/internal/services/http"
	"task_2/internal/services/logger"
	"task_2/internal/services/reader"
	"task_2/internal/services/sum"
)

type App struct{}

func (app *App) Run(source string) {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Открытие файлов логов и результата
	logFile, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	outputFile, err := os.OpenFile(cfg.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open output file: %v", err)
	}
	defer outputFile.Close()

	logger := logger.NewLogger(logFile)

	// Каналы для асинхронной обработки
	dataChan := make(chan []int)
	statusChan := make(chan int)
	errChan := make(chan error)

	// Выбор ридера на основе источника данных
	var dataReader reader.Reader
	if source == "stdin" {
		dataReader = reader.NewStdinReader()
	} else {
		dataReader = reader.NewFileReader(cfg.InputFile)
	}

	// Используем sync.WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	wg.Add(2)

	// Запуск горутины для чтения данных
	go func() {
		defer wg.Done()
		data, err := dataReader.ReadData()
		if err != nil {
			errChan <- fmt.Errorf("failed to read data: %v", err)
			return
		}
		dataChan <- data
	}()

	// Запуск горутины для выполнения HTTP GET запроса
	go func() {
		defer wg.Done()
		httpChecker := http.NewSimpleHTTPChecker()
		status, err := httpChecker.CheckURL(cfg.URL)
		if err != nil {
			errChan <- fmt.Errorf("failed to check URL: %v", err)
			return
		}
		statusChan <- status
	}()

	// Закрытие каналов по завершении всех горутин
	go func() {
		wg.Wait()
		close(dataChan)
		close(statusChan)
		close(errChan)
	}()

	// Обработка данных по мере их поступления
	for {
		select {
		case data, ok := <-dataChan:
			if ok {
				sumCalculator := sum.NewSimpleSumCalculator()
				totalSum := sumCalculator.SumNumbers(data)
				logger.LogInfo(fmt.Sprintf("Total sum: %d", totalSum))
				fmt.Printf("Total sum: %d\n", totalSum)
				if _, err := fmt.Fprintf(outputFile, "Total sum: %d\n", totalSum); err != nil {
					logger.LogError("Failed to write total sum to output file", err)
				}
			}
		case status, ok := <-statusChan:
			if ok {
				logger.LogInfo(fmt.Sprintf("URL status: %d", status))
				fmt.Printf("URL status: %d\n", status)
				if _, err := fmt.Fprintf(outputFile, "URL status: %d\n", status); err != nil {
					logger.LogError("Failed to write URL status to output file", err)
				}
			}
		case err := <-errChan:
			if err != nil {
				logger.LogError("Error occurred", err)
			}
			return
		}

		// Если все каналы закрыты, завершаем обработку
		if dataChan == nil && statusChan == nil && errChan == nil {
			break
		}
	}
}
