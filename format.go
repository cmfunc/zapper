package zapper

import (
	"go.uber.org/zap"
)

// Debug zapper对外暴露的日志打印方法
// fields: 为zap自定义的字段类型
// 不推荐使用zap.Any类型
func Debug(msg string, fields ...zap.Field) {
	defaultLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	defaultLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	defaultLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	defaultLogger.Error(msg, fields...)
}

// DPanic logs are particularly important errors. In development the
// logger panics after writing the message.
func DPanic(msg string, fields ...zap.Field) {
	defaultLogger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	defaultLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	defaultLogger.Fatal(msg, fields...)
}
