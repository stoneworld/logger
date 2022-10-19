package logger

import (
	"context"
)

func Access(log NormalAccessLog) {
	accessLogger.Info(`access`, log.AccessFields()...)
}

var bizLog = NewLogger().WithTag(`biz`)

// WithData 给日志附加自定义数据，在日志中会创建data字段
func WithData(data interface{}) Logger {
	return bizLog.WithData(data)
}

// WithTag 给日志打标签，在日志中会创建tag字段
func WithTag(tag string) Logger {
	return bizLog.WithTag(tag)
}

// WithoutStack 当为error和warn级别日志时，是否不打印调用栈
func WithoutStack(without bool) Logger {
	return bizLog.WithoutStack(without)
}

// AddCallerSkip 日志的调用栈和caller跳过几层
func AddCallerSkip(skip int) Logger {
	return bizLog.AddCallerSkip(skip)
}

// Debug 调试日志，生产环境一般不采集
func Debug(ctx context.Context, args ...interface{}) {
	bizLog.AddCallerSkip(1).Debug(ctx, args...)
}

// DebugF 调试日志，生产环境一般不采集
func DebugF(ctx context.Context, template string, args ...interface{}) {
	bizLog.AddCallerSkip(1).DebugF(ctx, template, args...)
}

// Info 一般信息日志
func Info(ctx context.Context, args ...interface{}) {
	bizLog.AddCallerSkip(1).Info(ctx, args...)
}

// InfoF 一般信息日志
func InfoF(ctx context.Context, template string, args ...interface{}) {
	bizLog.AddCallerSkip(1).InfoF(ctx, template, args...)
}

// Warn 注意日志，一般用于记录不影响主流程可以降级处理的问题
func Warn(ctx context.Context, args ...interface{}) {
	bizLog.AddCallerSkip(1).Warn(ctx, args...)
}

// WarnF 注意日志，一般用于记录不影响主流程可以降级处理的问题
func WarnF(ctx context.Context, template string, args ...interface{}) {
	bizLog.AddCallerSkip(1).WarnF(ctx, template, args...)
}

// Error 错误日志，一般用于导致程序无法继续进行的错误
func Error(ctx context.Context, args ...interface{}) {
	bizLog.AddCallerSkip(1).Error(ctx, args...)
}

// ErrorF 错误日志，一般用于导致程序无法继续进行的错误
func ErrorF(ctx context.Context, template string, args ...interface{}) {
	bizLog.AddCallerSkip(1).ErrorF(ctx, template, args...)
}

// Panic 恐慌日志，可能导致程序崩溃问题，会触发panic，不recover会导致程序崩溃
func Panic(ctx context.Context, args ...interface{}) {
	bizLog.AddCallerSkip(1).Panic(ctx, args...)
}
