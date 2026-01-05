package infra

import ( 
	"log"
	"os" 
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)
 

type AppLogger struct {
	*logrus.Logger
}

func NewLogger() *AppLogger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	})
	logger.SetLevel(logrus.InfoLevel)
	return &AppLogger{Logger: logger}
}

type DBLogger struct {
	logger.Interface
}

func NewDBLogger() *DBLogger {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 1 * time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	return &DBLogger{Interface: dbLogger}
}
