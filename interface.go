package logger

import (
	"context"
	"go.uber.org/zap"
)

type BasicLog interface {
	BasicFields() []zap.Field // 日志中的基础字段
	GetStackByError(err error) []byte
}

type NormalAccessLog interface {
	BasicLog
	GetTraceId() string
	AccessFields() []zap.Field // Access日志中的字段
	LogAddNotice(key string, v interface{})
	LogAddNoticeTime(key string, v interface{})
	TimeCost() func(funcName string)
}

type Logger interface {
	// WithData 给日志附加自定义数据，在日志中会创建data字段
	WithData(data interface{}) Logger
	// WithTag 给日志打标签，在日志中会创建tag字段
	WithTag(tag string) Logger
	// WithoutStack 当为error和warn级别日志时，是否不打印调用栈
	WithoutStack(without bool) Logger
	// AddCallerSkip 日志的调用栈和caller跳过几层
	AddCallerSkip(skip int) Logger
	// Debug 调试日志，生产环境一般不采集
	Debug(ctx context.Context, args ...interface{})
	// DebugF 调试日志，生产环境一般不采集
	DebugF(ctx context.Context, template string, args ...interface{})
	// Info 一般信息日志
	Info(ctx context.Context, args ...interface{})
	// InfoF 一般信息日志
	InfoF(ctx context.Context, template string, args ...interface{})
	// Warn 注意日志，一般用于记录不影响主流程可以降级处理的问题
	Warn(ctx context.Context, args ...interface{})
	// WarnF 注意日志，一般用于记录不影响主流程可以降级处理的问题
	WarnF(ctx context.Context, template string, args ...interface{})
	// Error 错误日志，一般用于导致程序无法继续进行的错误
	Error(ctx context.Context, args ...interface{})
	// ErrorF 错误日志，一般用于导致程序无法继续进行的错误
	ErrorF(ctx context.Context, template string, args ...interface{})
	// Panic 恐慌日志，可能导致程序崩溃问题，会触发panic，不recover会导致程序崩溃
	Panic(ctx context.Context, args ...interface{})
}
