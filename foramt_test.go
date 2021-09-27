package zapper

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDebug(t *testing.T) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
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
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	SetDefaultLogger(logger)

	defer defaultLogger.Sync()
	Info("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestWarn(t *testing.T) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	SetDefaultLogger(logger)

	defer defaultLogger.Sync()
	Warn("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
func TestDPanic(t *testing.T) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	SetDefaultLogger(logger)

	defer defaultLogger.Sync()
	DPanic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
}
// func TestPanic(t *testing.T) {
// 	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
// 	defer defaultLogger.Sync()
// 	Panic("benchmark zap ",
// 		zap.Int("int", 10),
// 		zap.Error(errors.New("text string")),
// 		zap.String("key string", "val string"),
// 		zap.Time("time", time.Now()),
// 	)
// }
// func TestFatal(t *testing.T) {
// 	defaultLogger = NewAdvancedLogger(zap.DebugLevel, "product", "module", "./zaper.log",time.Second*3)
// 	defer defaultLogger.Sync()
// 	Fatal("benchmark zap ",
// 		zap.Int("int", 10),
// 		zap.Error(errors.New("text string")),
// 		zap.String("key string", "val string"),
// 		zap.Time("time", time.Now()),
// 	)
// }

func TestSetDefaultLogger(t *testing.T) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
	defer logger.Sync()
	SetDefaultLogger(logger)

	Info("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)
	
}

func TestAdvancedLogger(t *testing.T) {
	syncCycle := time.Hour
	wr := NewFileWriter("./test.log", syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	logger := NewAdvancedLogger(zap.DebugLevel, "product", "module", syncer)
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
	DPanic("benchmark zap ",
		zap.Int("int", 10),
		zap.Error(errors.New("text string")),
		zap.String("key string", "val string"),
		zap.Time("time", time.Now()),
	)

}