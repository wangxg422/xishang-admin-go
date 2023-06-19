package logger

import (
	"backend/global"

	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	global.Log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	global.Log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	global.Log.Error(msg, fields...)
}
