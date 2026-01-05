package main

import (
	"context"
	"rest-fiber/config"
	"rest-fiber/internal/category"
	"rest-fiber/internal/infra"
	"rest-fiber/internal/post"
	"rest-fiber/internal/user"
)

func InitMigrate(ctx context.Context) error {
	dbLogger := infra.NewDBLogger()
	env, err := config.GetEnvs()
	if err != nil {
		return err
	}
	db, err := infra.GetDatabaseStandalone(env, dbLogger)
	if err != nil {
		return err
	}

	if err = db.Debug().Exec(`
		DO $$ BEGIN
			CREATE TYPE role_type AS ENUM ('ADMIN', 'MEMBER');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return err
	}

	if err = db.Debug().Exec(`
		DO $$ BEGIN
			CREATE TYPE status_type AS ENUM ('PUBLISHED', 'DRAFT');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`).Error; err != nil {
		return err
	}

	return db.WithContext(ctx).Debug().AutoMigrate(
		&user.User{},
		&post.Post{},
		&category.Category{},
	)
}
