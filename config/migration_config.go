package config

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type DB_CONFIG logger.Interface

func MigrationConfig() (string, DB_CONFIG) {
	var dsn string = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
	var logConfig DB_CONFIG = logger.New(
		writer,
		logger.Config{
			SlowThreshold:             1 * time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	return dsn, logConfig
}
