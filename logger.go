package zapper

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLevel    = zap.DebugLevel
	defaultProduct  = "zapper"
	defaultModule   = "zapper"
	defaultFilePath = "./zapper.log"
)

var defaultLogger *zap.Logger = NewLogger(defaultLevel, defaultProduct, defaultModule, defaultFilePath)

// 初始化函数
func Zapper(level zapcore.Level, product, module string, outputPath string) {
	logger := NewLogger(level, product, module, outputPath)
	SetDefaultLogger(logger)
}

func Sync() { defaultLogger.Sync() }

func SetDefaultLogger(logger *zap.Logger) { defaultLogger = logger }

// NewLogger 基于NewAdvancedLogger 的日志工厂方法，默认按整点自动切割日志文件
func NewLogger(level zapcore.Level, product, module string, outputPath string) *zap.Logger {

	syncCycle := time.Hour
	wr := NewWriter(outputPath, syncCycle, 0)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	return NewAdvancedLogger(level, product, module, syncer)
}

// NewAdvancedLogger 高级配置方法，可自动切割日志
func NewAdvancedLogger(level zapcore.Level, product, module string, syncer zapcore.WriteSyncer) *zap.Logger {

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= level })

	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "timestamp",
		NameKey:          "zapper",
		CallerKey:        "caller",
		FunctionKey:      "func",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.EpochTimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "\t",
	})

	core := zapcore.NewCore(fileEncoder, syncer, priority)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(level)).With(
		zap.String("product", product),
		zap.String("module", module),
	)

	return logger
}
