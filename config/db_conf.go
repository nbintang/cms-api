package config

import (
	"context"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(lc fx.Lifecycle, env Env, logger logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(env.DatabaseURL), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return sqlDB.Close()
		},
		OnStart: func(ctx context.Context) error {
			return nil
		},
	})
	
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}
