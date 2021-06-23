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
	"os"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type EnvType int

const (
	EnvLocal    EnvType = iota //本地环境
	EnvLiantiao                //联调环境
	EnvQA                      //QA环境
	EnvPro                     //生产环境
)

// Config 日志包配置
// TODO 注意结构体字段，内存对其
type Config struct {
	Env       EnvType       //int 当前环境
	FileLevel zapcore.Level //int8 日志输出等级
	LogFile   string        //string 日志文件位置
}

// NewConfig 初始化一个日志配置对象
func NewConfig(loglevel zapcore.Level) *Config {
	return &Config{
		FileLevel: loglevel,
	}
}

// logger only object for format log
var logger *zap.Logger

// aviod mutil to init logger
var once sync.Once

// zaper package global config object
var zaperCfg *Config

// newFileSyncer 构造一个文件同步器
func newFileSyncer() io.Writer {
	rl, err := rotatelogs.New(
		zaperCfg.LogFile+".%Y%m%d%H%M",
		rotatelogs.WithClock(rotatelogs.UTC),
		rotatelogs.WithLinkName(zaperCfg.LogFile),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(-1),
	)
	if err != nil {
		panic(err)
	}
	return rl
}

// Init 初始化zaper包
func Init(cfg *Config) {
	once.Do(func() {
		// init zapperCfg var
		zaperCfg = cfg
		// First, define our level-handling logic.
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zap.ErrorLevel
		})

		filePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zaperCfg.FileLevel
		})

		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		topicErrors := zapcore.AddSync(newFileSyncer())

		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		cores := make([]zapcore.Core, 0)
		cores = append(cores, zapcore.NewCore(fileEncoder, topicErrors, filePriority))
		if zaperCfg.Env == EnvLocal {
			cores = append(cores,
				zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
				zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
		}

		core := zapcore.NewTee(cores...)

		logger = zap.New(core)
		// defer logger.Sync() 进程退出前调用

		logger.Info("init zapper success", zap.Int8("return", 0))

	})
}
