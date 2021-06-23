// zaper
// 暴露给用户使用的方法
// 防止日志库 panic defer recover
package zaper

import (
	"context"

	"go.uber.org/zap"
)

func format(ctx context.Context,msg string, fields ...zap.Field)  {
	headerFields:=getLogHeaderFromCtx(ctx).fieldsMarshal()
	fields=append(fields, headerFields...)
}

func Debug(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Debug(msg , fields...)
}

func Info(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Info(msg , fields...)
}

func Warn(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Warn(msg , fields...)
}

func Error(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Error(msg , fields...)
}

func DPanic(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.DPanic(msg , fields...)
}

func Panic(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Panic(msg , fields...)
}

// Fatal 写入日志以后，直接调用os.Exit(1)
func Fatal(ctx context.Context,msg string, fields ...zap.Field)  {
	defer logcatch()
	format(ctx,msg,fields...)
	logger.Fatal(msg , fields...)
}
