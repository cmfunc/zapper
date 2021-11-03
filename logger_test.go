package zapper

import (
	"errors"
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapperRotate(t *testing.T) {
	syncCycle := time.Minute
	wr := NewWriter("./test_default_filed.log", syncCycle, 0)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	SetDefaultLogger(logger)
	Sync()
	for i := 0; i < 360; i++ {
		Error("msg string",
			zap.Int("i", i),
			zap.String("msg", "啥也不是"),
			zap.Int("i", 123),
		)
		time.Sleep(time.Second)
	}
}


func TestZapper(t *testing.T) {
	Zapper(zap.DebugLevel, "product", "module", "./test.log")

	var wg = &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			Error("benchmark zap ",
				zap.Int("int", 10),
				zap.Error(errors.New("text string")),
				zap.String("key string", "val string"),
				zap.Time("time", time.Now()))
			wg.Done()
		}(wg)
	}
	wg.Wait()
	defer defaultLogger.Sync()

}

func TestNewAdvancedLogger(t *testing.T) {
	syncCycle := time.Second
	wr := NewWriter("./test.log", syncCycle, 0)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	defer logger.Sync()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 3)
		logger.Error("msg string", zap.Int("i", i))
	}
}

func BenchmarkZapperInfo(b *testing.B) {
	syncCycle := time.Hour
	wr := NewWriter("./test.log", syncCycle, 0)
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
