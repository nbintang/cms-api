package main

import (
	"rest-fiber/config"
	"rest-fiber/internal"
	"rest-fiber/internal/auth"
	"rest-fiber/internal/user"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}
	fx.New(
		config.Module,
		internal.Module,
		user.Module,
		auth.Module,
	).Run()
}
