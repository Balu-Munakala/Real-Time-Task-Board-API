package logger

import (
	"go.uber.org/zap"
)

func NewLogger(level string) *zap.Logger {
	if level == "development" {
		return zap.Must(zap.NewDevelopment(zap.AddCallerSkip(8)))
	}
	return zap.Must(zap.NewProduction(zap.AddCallerSkip(-10)))
}
