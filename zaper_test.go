package zaper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestAdvancedLoggerDebug(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log")
	defer logger.Sync()

	logger.Debug("msg string", zap.Int("int", 10))

	SetDefaultLogger(logger)
	logger.Warn("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)

}

func TestAdvancedLoggerDefault(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log")
	defer logger.Sync()
	Debug("msg string", zap.Int("int", 10))
	Warn("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	Error("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	Info("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	DPanic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	Panic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	Fatal("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}

func BenchmarkZaperInfo(b *testing.B) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./test.log")
	defer logger.Sync()

	for i := 0; i < b.N; i++ {
		logger.Error("benchmark zap ",
			zap.Int("int", 10),
			zap.Error(errors.New("text string")),
			zap.String("key string", "val string"),
			zap.Time("time", time.Now()),
		)
	}

}
