package zaper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestAdvanceaper(t *testing.T) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log")
	defer logger.Sync()

	logger.Debug("msg string", zap.Int("int", 10))

}

func BenchmarkZaperInfo(b *testing.B) {
	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log")
	defer logger.Sync()

	for i := 0; i < b.N; i++ {
		logger.Error("benchmark zap ",
			zap.Int("int", 10),
			zap.Error(errors.New("text string")),
			zap.String("key string", "val string"),
			zap.Time("timr", time.Now()),
		)
	}

}
