package main

import (
	"context"
	"log"
	"rest-fiber/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMigrate(ctx context.Context) error {
	dsn, logConfig := config.MigrationConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logConfig,
	})
	if err != nil {
		return err
	}
	err = db.WithContext(ctx).AutoMigrate(
	//Entry models
	// ...
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	if err := InitMigrate(ctx); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
