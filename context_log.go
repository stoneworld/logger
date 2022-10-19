package logger

import (
	"context"
)

func GetTraceId(ctx context.Context) string {
	return ToAccessLog(ctx).GetTraceId()
}

func AccessTimeCost(ctx context.Context) func(funcName string) {
	return ToAccessLog(ctx).TimeCost()
}

func AddAccessNotice(ctx context.Context, key string, v interface{}) {
	ToAccessLog(ctx).LogAddNotice(key, v)
}

//goland:noinspection GoUnusedExportedFunction
func AddAccessNoticeTime(ctx context.Context, key string, v interface{}) {
	ToAccessLog(ctx).LogAddNoticeTime(key, v)
}

func ToAccessLog(ctx context.Context) NormalAccessLog {
	if ctx == nil {
		return nil
	}
	accessLog, ok := ctx.Value(`accessLog`).(NormalAccessLog)
	if ok {
		return accessLog
	}
	return nil
}

func ToBasicLog(ctx context.Context) BasicLog {
	if ctx == nil {
		return nil
	}
	basicLog, ok := ctx.Value(`accessLog`).(BasicLog)
	if ok {
		return basicLog
	}
	return nil
}
