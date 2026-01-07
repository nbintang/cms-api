package infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-fiber/config"
)

func GetDatabaseStandalone(env config.Env, logger *DBLogger) (*gorm.DB, error) {
	sslMode := env.DatabaseSSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.DatabaseHost,
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseName,
		env.DatabasePort,
		sslMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
}
func NewDatabase(env config.Env, logger *DBLogger) (*gorm.DB, error) {
	return GetDatabaseStandalone(env, logger)
}
