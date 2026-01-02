package main

import (
	"context"
	"rest-fiber/config"
	"rest-fiber/internal/category"
	"rest-fiber/internal/post"
	"rest-fiber/internal/user"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
		&user.User{},
		&post.Post{},
		&category.Category{},
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}
	viper.AutomaticEnv()
	if err = InitMigrate(ctx); err != nil {
		logrus.Fatalf("Migration failed: %v", err)
	}
	logrus.Println("Migration Succeed")
}
