// zaper
// 禁止使用zap.Suger
// 尽量避免使用反射做类型转换
// 变量都使用明确的类型
// 使用本包时，需要自行实现日志的Config配置，调用NewConfig()方法，将配置注入本包
// 然后调用Init()函数，对日志对象logger进行初始化
// 完成以后即可调用日志输出方法：Error()等
package zaper

import (
	"io"
	"io/ioutil"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 日志包配置
// TODO 注意结构体字段，内存对其
type Config struct {
	//int8 日志输出等级
	HandlingLevel zapcore.Level
	//string 日志文件位置
	LogFile string
	// 日志落盘writer
	ConsoleWriter io.Writer
	FileWriter    io.Writer
	QueueWriter   io.Writer
	// 日志编码器Encoder
	// TODO 颜色编码
	ConsoleEncoder *zapcore.EncoderConfig
	FileEncoder    *zapcore.EncoderConfig
	QueueEncoder   *zapcore.EncoderConfig
}

// 日志落盘类型
type LogWriter func(p []byte) (n int, err error)

// 实现io.Writer接口
func (l LogWriter) Write(p []byte) (n int, err error) {
	return l(p)
}

// NewConfig 初始化一个日志配置对象
func NewConfig(loglevel zapcore.Level) *Config {
	return &Config{
		HandlingLevel: loglevel,
	}
}

// logger only object for format log
var logger *zap.Logger

// aviod mutil to init logger
var once sync.Once

// zaper package global config object
var zapperCfg *Config

// Init 初始化zaper包
func Init(cfg *Config) {
	once.Do(func() {
		// init zapperCfg var
		zapperCfg = cfg
		// First, define our level-handling logic.
		// 通过日志等级来切分日志
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapperCfg.HandlingLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapperCfg.HandlingLevel
		})

		core := zapcore.NewTee(
			initConsoleCore(highPriority),
			initConsoleCore(lowPriority),
			initFileCore(highPriority),
			initFileCore(lowPriority),
			initQueueCore(highPriority),
			initQueueCore(lowPriority),
		)

		logger = zap.New(core)
		defer logger.Sync()

		logger.Info("init zapper success", zap.Int8("return", 0))

	})
}

// initConsoleCore 初始化控制台core
// 日志输出到控制台
func initConsoleCore(priority zap.LevelEnablerFunc) zapcore.Core {
	var writer io.Writer // TODO attain a writer
	var encoderCfg zapcore.EncoderConfig

	writer = zapperCfg.ConsoleWriter
	if writer == nil {
		writer = os.Stdout
	}

	if zapperCfg.ConsoleEncoder == nil {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = *zapperCfg.ConsoleEncoder
	}
	return initOneCore(priority, writer, encoderCfg)
}

// initFileCore 初始化文件core
// 日志落到文件中
// TODO: 日志切割、按时间间隔、按文件大小、
func initFileCore(priority zap.LevelEnablerFunc) zapcore.Core {
	var writer io.Writer // TODO attain a writer
	var encoderCfg zapcore.EncoderConfig

	writer = zapperCfg.FileWriter
	if writer == nil {
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   zapperCfg.LogFile, // 日志文件所在位置
			MaxSize:    500,               // megabytes 日志文件最大限制
			MaxBackups: 0,                 // 最大保留历史日志数量，0默认保留所有
			MaxAge:     0,                 // days，日志最大保留天数，0默认保留所有
			Compress:   true,              // disabled by default
		})
	}

	if zapperCfg.FileEncoder == nil {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = *zapperCfg.FileEncoder
	}
	return initOneCore(priority, writer, encoderCfg)
}

// initQueueCore 初始化队列core
// 日志落到消息队列中
func initQueueCore(priority zap.LevelEnablerFunc) zapcore.Core {
	var writer io.Writer // TODO attain a writer
	var encoderCfg zapcore.EncoderConfig

	writer = zapperCfg.QueueWriter
	if writer == nil {
		writer = ioutil.Discard
	}

	if zapperCfg.QueueEncoder == nil {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = *zapperCfg.QueueEncoder
	}

	return initOneCore(priority, writer, encoderCfg)
}

func initOneCore(priority zap.LevelEnablerFunc, writer io.Writer, encoderCfg zapcore.EncoderConfig) zapcore.Core {
	// 日志落盘 实现io.Writer
	topicErrors := zapcore.AddSync(writer)

	// zap.NewProductionEncoderConfig 自定义配置
	kafkaEncoder := zapcore.NewJSONEncoder(encoderCfg)

	return zapcore.NewCore(kafkaEncoder, topicErrors, priority)

}
