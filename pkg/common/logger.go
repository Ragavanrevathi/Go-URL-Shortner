package common

import (
	"context"
	"go.uber.org/zap"
)

func LogWithTrace(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(LoggerKey).(*zap.Logger)
	if !ok {
		l, _ := zap.NewProduction()
		return l
	}
	return logger
}
