package config

import (
	"os"

	"github.com/sirupsen/logrus"
)


func NewAppLogger() *logrus.Logger {
	logger := logrus.New()
	logrus.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}