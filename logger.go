// zaper
// !!! forbidden to use zap.Suger
// main goroutine with graceful exit ,logger.Sync() and corn.Stop()
package zapper

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


const (
	defaultLevel    = zap.DebugLevel
	defaultProduct  = "zaper"
	defaultModule   = "zaper"
	defaultFilePath = "./zaper.log"
)

var defaultLogger *zap.Logger = NewBasicLogger(defaultLevel, defaultProduct, defaultModule, defaultFilePath)

func SetDefaultLogger(logger *zap.Logger) { defaultLogger = logger }



// NewLogger 基于NewAdvancedLogger 的日志工厂方法，默认按整点自动切割日志文件
func NewLogger(level zapcore.Level, product, module string, outputPath string) *zap.Logger  {

	syncCycle := time.Hour
	wr := NewFileWriter(outputPath, syncCycle)
	syncer := zapcore.AddSync(wr) //ioutil.Discard

	return NewAdvancedLogger(level, product, module, syncer)
}


// NewAdvancedLogger 高级配置方法，可自动切割日志
func NewAdvancedLogger(level zapcore.Level, product, module string,syncer zapcore.WriteSyncer) *zap.Logger {

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= level })

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

	return logger
}


// NewBasicLogger
// 不建议使用，不能自动切割日志
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