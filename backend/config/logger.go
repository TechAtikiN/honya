package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

const logDir = "./logs"

var logFile *os.File

func removeOlderLogs() {
	_ = godotenv.Load()

	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Printf("Error reading logs directory: %v", err)
		return
	}

	LOG_RETENTION, err := strconv.Atoi(os.Getenv("LOG_RETENTION"))
	if err != nil {
		LOG_RETENTION = 7 // Default to 7 days
	}

	cutoffDate := time.Now().AddDate(0, 0, -LOG_RETENTION)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileDate, err := time.Parse("2006-01-02", file.Name()[:10])
		if err != nil {
			continue
		}

		if fileDate.Before(cutoffDate) {
			filePath := filepath.Join(logDir, file.Name())
			if err := os.Remove(filePath); err != nil {
				log.Printf("Error deleting old log file %s: %v", file.Name(), err)
			} else {
				log.Printf("Deleted old log file: %s", file.Name())
			}
		}
	}
}

func getLogFile() (*os.File, error) {
	_ = godotenv.Load()
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	removeOlderLogs()

	var logPath string
	if os.Getenv("LOG_STACK") == "daily" {
		currentTime := time.Now()
		fileName := fmt.Sprintf("%s.log", currentTime.Format("2006-01-02"))
		logPath = filepath.Join(logDir, fileName)
	} else {
		logPath = filepath.Join(logDir, "error.log")
	}

	return os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func SetupLogger() fiber.Handler {
	var err error
	logFile, err = getLogFile()
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	return logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Format:     "\n" + `{"timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}", "error":"${error}", "response":[${responseBody}], "ip":"${ip}"}` + "\n",
		Output:     logFile,
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() >= 400 {
				_, err := logFile.Write(logString)
				if err != nil {
					log.Printf("Error writing to log file: %v", err)
				}
			}
		},
	})
}

func CloseLogFile() {
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			log.Printf("Error closing log file: %v", err)
		}
	}
}
