package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

func MigrationConfig() (string, logger.Interface) {
	var dsn string = viper.GetString("DATABASE_URL")
	writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
	var logConfig logger.Interface = logger.New(
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
