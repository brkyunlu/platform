package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"sync"
)

// LogLevel represents log levels
type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Error
	// ... Other log levels can be added here
)

var (
	log      *logrus.Logger
	logMutex sync.Mutex
)

// InitLogger initializes the logger with default configurations
func InitLogger() {
	logMutex.Lock()
	defer logMutex.Unlock()
	log = logrus.New()
	if log != nil {
		return
	}

	// Set log levels
	log.SetLevel(logrus.InfoLevel) // Use logrus.DebugLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel as needed

	// Set log format
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// Set up logging to specific files and rotation
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  "info.log",
		logrus.ErrorLevel: "error.log",
	}, &logrus.JSONFormatter{})) // Save in JSON format

}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

// SetLogLevel sets the log level
func SetLogLevel(level LogLevel) {
	if log == nil {
		InitLogger()
	}

	switch level {
	case Info:
		log.SetLevel(logrus.InfoLevel)
	case Warn:
		log.SetLevel(logrus.WarnLevel)
	case Error:
		log.SetLevel(logrus.ErrorLevel)
		// ... Other log levels can be added here
	}
}
