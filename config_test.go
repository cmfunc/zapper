package zaper

import (
	"testing"
	"time"

	"git.ymt360.com/go/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapInfo(t *testing.T) {
	cfg := NewConfig(zapcore.DebugLevel, "./zap_benchmark.log")
	Init(cfg)
	logger.Info("msg string", zap.String("key string", "testing"))
}

func BenchmarkZapInfo(b *testing.B) {
	cfg := NewConfig(zapcore.DebugLevel, "./zap_benchmark.log")
	Init(cfg)
	for i := 0; i < b.N; i++ {
		logger.Info("msg string", zap.String("key string", "testing"))
	}
}

func BenchmarkCombinationParallelZapInfo(b *testing.B) {
    cfg := NewConfig(zapcore.DebugLevel, "./zap_benchmark.log")
	Init(cfg)
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            logger.Info("msg string", zap.String("key string", "testing"))
        }
    })
}

func TestLoggingInfo(t *testing.T) {
	msg := map[string]string{"AAAA": "BBBBB"}
	c := logging.LogConfig{Path: "./", File: "logging_benchmark.log", Mode: 1, Rotate: false, Debug: false}
	logger := logging.NewLogger(&c)
	header := logging.LogHeader{LogId: "abc"}

	logger.Info(&header, msg)

	time.Sleep(100 * time.Millisecond)
}

func BenchmarkLoggingInfo(b *testing.B) {
	msg := map[string]string{"AAAA": "BBBBB"}
	c := logging.LogConfig{Path: "./", File: "logging_benchmark.log", Mode: 1, Rotate: false, Debug: false}
	logger := logging.NewLogger(&c)
	header := logging.LogHeader{LogId: "abc"}

	for i := 0; i < b.N; i++ {
		logger.Info(&header, msg)
	}

	time.Sleep(100 * time.Millisecond)

}
