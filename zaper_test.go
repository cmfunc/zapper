package zaper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	SetDefaultLogger(logger)
	defer defaultLogger.Sync()
	Debug("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestInfo(t *testing.T) {
	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer defaultLogger.Sync()
	Info("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestWarn(t *testing.T) {
	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer defaultLogger.Sync()
	Warn("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestDPanic(t *testing.T) {
	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer defaultLogger.Sync()
	DPanic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestPanic(t *testing.T) {
	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer defaultLogger.Sync()
	Panic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestFatal(t *testing.T) {
	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer defaultLogger.Sync()
	Fatal("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}

func TestAdvancedLoggerInfo(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer logger.Sync()
	SetDefaultLogger(logger)
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

func TestAdvancedDefaultLogger(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer logger.Sync()

	SetDefaultLogger(logger)
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

func TestAdvancedLogger(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
	defer logger.Sync()
	logger.Debug("msg string", zap.Int("int", 10))
	logger.Warn("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	logger.Error("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	logger.Info("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	logger.DPanic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	logger.Panic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	logger.Fatal("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}

func BenchmarkZaperInfo(b *testing.B) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./test.log",time.Hour*1)
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
