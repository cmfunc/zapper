// zaper
// !!! forbidden use zap.Suger
// main goroutine with graceful exit ,logger.Sync() and corn.Stop()
package zaper

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewBasicLogger方法
func NewBasicLogger(level zapcore.Level, product, module string, outputPath string) *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "timestamp",
			NameKey:          "zaper",
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
		},
		OutputPaths:      []string{outputPath},
		ErrorOutputPaths: []string{outputPath},
		InitialFields: map[string]interface{}{
			"product": product,
			"module":  module,
		},
	}

	logger, err := cfg.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}

	return logger
}

// NewAdvancedLogger 高级配置方法，可自动切割日志
func NewAdvancedLogger(level zapcore.Level, product, module string, outputPath string, syncCycle time.Duration) *zap.Logger {

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= level })

	wr := NewFileWriter(outputPath, syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "timestamp",
		NameKey:          "zaper",
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
	// defer logger.Sync()
	return logger
}

const (
	defaultLevel    = zap.DebugLevel
	defaultProduct  = "zaper"
	defaultModule   = "zaper"
	defaultFilePath = "./zaper.log"
)

var defaultLogger *zap.Logger = NewBasicLogger(defaultLevel, defaultProduct, defaultModule, defaultFilePath)

func SetDefaultLogger(logger *zap.Logger) { defaultLogger = logger }

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
