package configuration

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitialiseLogger(l *logrus.Logger) {
	logger = l
}

func LogInfo(message string) {
	if logger != nil {
		logger.Info(message)
	}
}

func LogError(message string) {
	if logger != nil {
		logger.Error(message)
	}
}
