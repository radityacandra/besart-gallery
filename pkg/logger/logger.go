package logger

import (
	"go.uber.org/zap"
)

func LoadLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	logger = logger.With(zap.Any("service", "job-portal"))

	return logger, err
}
