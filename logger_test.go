package zapper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkZaperInfo(b *testing.B) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
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
