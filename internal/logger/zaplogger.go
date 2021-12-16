package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	MyLogger *zap.SugaredLogger
}

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
