package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

var accessLogger *zap.Logger

var level zapcore.Level

type Config struct {
	LogLevel string
	// 是否开启debug模式，未开启debug模式，仅记录错误
	Debug bool
	// 日志输出的方式
	// none为不输出日志，file 为文件方式输出，console为控制台。默认为console
	Output string
	// 日志文件路径
	File string
	// 日志文件的保留时间单位是天，超过的将会被删除
	MaxAge  int
	AppName string
}

func InitLogger(config *Config) {
	switch strings.ToUpper(config.LogLevel) {
	case "ERROR":
		level = zap.ErrorLevel
	case "WARN", "WARNING":
		level = zap.WarnLevel
	case "INFO":
		level = zap.InfoLevel
	case "DEBUG":
		level = zap.DebugLevel
	default:
		level = zap.DebugLevel
	}

	// 一般日志（info/warn/error）
	initCommonLogger(config)

	// access日志
	initAccessLogger(config)
}

func initCommonLogger(config *Config) {
	//debug 模式直接打印，部署到服务器时写文件
	if config.Debug {
		initConsoleCommonLogger(config)
	} else {
		initFileCommonLogger(config)
	}
}

func initAccessLogger(config *Config) {
	pathTemplate := config.File + config.AppName + `_%Y%m%d_access.log`

	accessLogger = NewZapLogger(pathTemplate, zap.InfoLevel, nil, config,
		zap.WithCaller(false),
		zap.Fields(zap.String(`tag`, `lib.access`)),
	)
}

func initFileCommonLogger(config *Config) {
	basePath := config.File + config.AppName
	info := NewZapLogger(basePath+"_%Y%m%d_info.log", level, nil, config, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	err := NewZapLogger(basePath+"_%Y%m%d_error.log", level, nil, config, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))

	stdLog, _ := zap.NewDevelopment()
	stdLogger, _ := zap.NewStdLogAt(stdLog, zapcore.DebugLevel)
	log.SetOutput(stdLogger.Writer())

	debugLogger = info
	infoLogger = info
	warnLogger = err
	errorLogger = err
	panicLogger = err
}

func initConsoleCommonLogger(config *Config) {
	devLogger := NewZapLogger(``, level, nil, config)
	stdLogger, _ := zap.NewStdLogAt(devLogger, zapcore.DebugLevel)
	log.SetOutput(stdLogger.Writer())
	debugLogger = devLogger
	infoLogger = devLogger
	warnLogger = devLogger
	errorLogger = devLogger
	panicLogger = devLogger
}
