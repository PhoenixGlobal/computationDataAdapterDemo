package utils

import (
	"computationDataAdapterDemo/conf"
	"go.uber.org/zap"
)

var MainLogger=NewLogger("Main.log", conf.Conf.ZapLog.LogLevel, true, true)

func GetLogger(loggerName string) *zap.Logger {
	var Logger *zap.Logger
	Logger = NewLogger(loggerName+".log", conf.Conf.ZapLog.LogLevel, true, true)
	return Logger
}