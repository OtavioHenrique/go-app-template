package logger

import "go.uber.org/zap"

// This class should containg all logic related to logic. If needed it can receive config or other port needed
func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger.Sugar()
}
