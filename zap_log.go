package logger

import (
	jsoniter "github.com/json-iterator/go"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"time"
)

func NewZapLogger(pathTemplate string, level zapcore.Level, encoderConfig *zapcore.EncoderConfig, config *Config, options ...zap.Option) *zap.Logger {
	if encoderConfig == nil {
		encoderConfig = getDefaultEncoderConfig(pathTemplate)
	}
	if pathTemplate != `` {
		encoder := zapcore.NewJSONEncoder(*encoderConfig)
		writer := zapcore.AddSync(getWriter(pathTemplate, config))
		return zap.New(zapcore.NewCore(encoder, writer, level), options...)
	} else {
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig = *encoderConfig
		zapConfig.Level = zap.NewAtomicLevelAt(level)
		devLogger, _ := zapConfig.Build(options...)
		return devLogger
	}
}

func getWriter(filename string, config *Config) io.Writer {
	hook, err := rotatelogs.New(
		filename,
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(config.MaxAge)),
	)
	if err != nil {
		log.Panic(err)
	}
	return hook
}

func getDefaultEncoderConfig(pathTemplate string) (config *zapcore.EncoderConfig) {
	var encodeLevel = zapcore.CapitalColorLevelEncoder

	if pathTemplate != `` {
		encodeLevel = zapcore.LowercaseLevelEncoder
	}
	var encoderConfig = zapcore.EncoderConfig{
		TimeKey:       "dateTime",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   encodeLevel,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(`2006-01-02 15:04:05.999`))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		NewReflectedEncoder: func(writer io.Writer) zapcore.ReflectedEncoder {
			return jsoniter.ConfigFastest.NewEncoder(writer)
		},
	}
	return &encoderConfig
}
