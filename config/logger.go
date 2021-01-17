package config

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var initLogOnce sync.Once

// GetLogger creates an instance of logrus.Logger depending on the environment and returns it
func GetLogger() *logrus.Logger {
	initLogOnce.Do(initLogger) // executes only sessionOnce
	return logger
}

func initLogger() {
	if logger == nil {
		panic("Call SetupLogger before any GetLogger calls")
	}

}

func SetupLogger() {
	newLogger := logrus.New()

	if !GetAppConfig().Production {
		newLogger.SetOutput(os.Stdout)
		newLogger.SetLevel(logrus.TraceLevel)
		newLogger.SetReportCaller(true)
	} else {
		newLogger.SetOutput(os.Stderr)
		newLogger.SetLevel(logrus.InfoLevel)
		newLogger.SetFormatter(&logrus.JSONFormatter{})
	}

	logger = newLogger
}
