package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	// Create a new logger instance
	Logger = logrus.New()

	// Set log output (to stdout in this case)
	Logger.SetOutput(os.Stdout)

	// Set log format (e.g., JSONFormatter or TextFormatter)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level (e.g., Info, Debug, Warn)
	Logger.SetLevel(logrus.InfoLevel)
}

