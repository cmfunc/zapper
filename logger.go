package zapper

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// defaultLogger 默认的日志对象
var defaultLogger *zap.Logger

// 初始化函数
// level 日志收集的最低级别
// product 产品线 例如uc
// module 模块名 例如user、passport、organization
// outputPath 日志文件地址,生成的日志文件会拼接整点时间，作区分
// Zapper是最方便的初始化zapper库的方法，初始化完成后，可以直接调用zapper.Info()、zapper.Error()打印日志
func Zapper(level zapcore.Level, product, module string, outputPath string) {
	logger := NewLogger(level, product, module, outputPath)
	SetDefaultLogger(logger)
}

// Sync 在服务停止前调用,日志罗盘goroutine生命周期管理的收尾工作
// 通知writer 及时停止；
// defer zapper.Sync()
func Sync() {
	defaultLogger.Sync()
	defaultWriter.Close()
}

// SetDefaultLogger
// logger 用于替换zapper中默认的logger
func SetDefaultLogger(logger *zap.Logger) { defaultLogger = logger }

// NewLogger 基于NewAdvancedLogger 的日志工厂方法，默认按整点自动切割日志文件
// level 日志收集的最低级别
// product 产品线 例如uc
// module 模块名 例如user、passport、organization
// outputPath 日志文件地址,生成的日志文件会拼接整点时间，作区分
// 返回一个新的Logger对象
func NewLogger(level zapcore.Level, product, module string, outputPath string) *zap.Logger {

	syncCycle := time.Hour
	wr := NewWriter(outputPath, syncCycle, 0)
	SetDefaultWriter(wr)

	syncer := zapcore.AddSync(wr) //ioutil.Discard

	return NewAdvancedLogger(level, product, module, syncer)
}

// NewAdvancedLogger 高级配置方法，可自动切割日志
// level 日志收集的最低级别
// product 产品线 例如uc
// module 模块名 例如user、passport、organization
// syncer 日志罗盘的Writer
// 返回一个新的Logger对象
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
