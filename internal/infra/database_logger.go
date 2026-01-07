package infra

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type DBLogger struct {
	logger.Interface
}

func NewDBLogger() *DBLogger {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 1 * time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	return &DBLogger{Interface: dbLogger}
}
