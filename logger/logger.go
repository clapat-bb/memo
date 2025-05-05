package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("init log system failed: " + err.Error())
	}
	Log = logger.Sugar()
}
