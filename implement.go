package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

var warnLogger *zap.Logger
var panicLogger *zap.Logger
var debugLogger *zap.Logger
var errorLogger *zap.Logger
var infoLogger *zap.Logger

var _ Logger = &defaultLogger{}

type defaultLogger struct {
	data         interface{}
	tag          string
	alarm        bool
	skip         int
	withoutStack bool
}

func (l *defaultLogger) clone() *defaultLogger {
	cp := *l
	return &cp

}
func (l *defaultLogger) WithData(data interface{}) Logger {
	nl := l.clone()
	nl.data = data
	return nl
}

func (l *defaultLogger) WithTag(tag string) Logger {
	nl := l.clone()
	nl.tag = tag
	return nl
}

func (l *defaultLogger) WithAlarm(alarm bool) Logger {
	cp := l.clone()
	cp.alarm = alarm
	return cp
}

func (l *defaultLogger) WithoutStack(without bool) Logger {
	nl := l.clone()
	nl.withoutStack = without
	return nl
}

func (l *defaultLogger) AddCallerSkip(skip int) Logger {
	cp := l.clone()
	cp.skip = cp.skip + skip
	return cp
}

func (l *defaultLogger) Debug(ctx context.Context, args ...interface{}) {
	l.logger(ctx, debugLogger).Sugar().Debug(sprint(args...))
}

func (l *defaultLogger) DebugF(ctx context.Context, template string, args ...interface{}) {
	l.logger(ctx, debugLogger).Sugar().Debugf(template, args...)
}

func (l *defaultLogger) Info(ctx context.Context, args ...interface{}) {
	l.logger(ctx, infoLogger).Sugar().Info(sprint(args...))
}

func (l *defaultLogger) InfoF(ctx context.Context, template string, args ...interface{}) {
	l.logger(ctx, infoLogger).Sugar().Infof(template, args...)
}

func (l *defaultLogger) Warn(ctx context.Context, args ...interface{}) {
	l.logger(ctx, warnLogger).Warn(sprint(args...))
}

func (l *defaultLogger) WarnF(ctx context.Context, template string, args ...interface{}) {
	l.logger(ctx, warnLogger).Warn(fmt.Sprintf(template, args...))
}

func (l *defaultLogger) Error(ctx context.Context, args ...interface{}) {
	l.logger(ctx, errorLogger).Error(sprint(args...))
	l.sendAlarm(ctx, nil, "", args...)
}

func (l *defaultLogger) ErrorF(ctx context.Context, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.logger(ctx, panicLogger).Error(msg)
	l.sendAlarm(ctx, nil, "", msg)
}

func (l *defaultLogger) Panic(ctx context.Context, args ...interface{}) {
	l.logger(ctx, panicLogger).Panic(sprint(args...))
	l.sendAlarm(ctx, nil, "", args)
}

func (l *defaultLogger) logger(ctx context.Context, log *zap.Logger) (newLogger *zap.Logger) {
	newLogger = l.fieldsFromContext(ctx, log)
	if l.data != nil {
		newLogger = newLogger.With(
			zap.String(`tag`, l.tag),
			zap.Any(`data`, l.data),
		)
	} else {
		newLogger = newLogger.With(
			zap.String(`tag`, l.tag),
		)
	}
	return newLogger.WithOptions(zap.AddCallerSkip(1 + l.skip))
}

func (l *defaultLogger) fieldsFromContext(ctx context.Context, log *zap.Logger) *zap.Logger {
	basicLog := ToBasicLog(ctx)
	if basicLog != nil {
		return log.With(basicLog.BasicFields()...)
	}
	return log
}

func (l *defaultLogger) sendAlarm(ctx context.Context, err error, stack string, args ...interface{}) {

}

func sprint(args ...interface{}) string {
	if len(args) > 0 {
		if len(args) == 1 {
			if str, ok := args[0].(string); ok {
				return str
			}
		}
		s := fmt.Sprintln(args...)
		return s[:len(s)-1]
	}
	return ``
}

func NewLogger() Logger {
	return &defaultLogger{}
}
