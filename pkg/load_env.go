package pkg

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv( ) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}
}
