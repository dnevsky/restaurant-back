package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	Log = logger.Sugar()
}
