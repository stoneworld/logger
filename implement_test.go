package logger

import (
	"context"
	"fmt"
	"github.com/cornelk/hashmap"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestAccess(t *testing.T) {
	config := Config{
		LogLevel: "info",
		Debug:    true,
		Output:   "console",
		File:     "../logs/",
		MaxAge:   7,
		AppName:  "wtf",
	}
	InitLogger(&config)
	accessLog := &AccessLog{
		LogId:       uuid.NewString(),
		Host:        "host",
		ContentType: "contentType",
		RequestURI:  "requestURI",
		Method:      "method",
		BeginTime:   time.Now(),
		EndTime:     time.Now(),
		NoticeInfo:  hashmap.New(32),
		NoticeTime:  hashmap.New(32),
		Status:      0,
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "accessLog", accessLog)
	AddAccessNotice(ctx, "noticeInfo", accessLog)
	end := AccessTimeCost(ctx)
	time.Sleep(1 * time.Second)
	end("testFunTime")
	Access(accessLog)
}

func TestWarnF(t *testing.T) {
	accessLog := AccessLog{
		LogId:       uuid.NewString(),
		Host:        "host",
		ContentType: "contentType",
		RequestURI:  "requestURI",
		Method:      "method",
		BeginTime:   time.Now(),
		EndTime:     time.Now(),
		NoticeInfo:  hashmap.New(32),
		NoticeTime:  hashmap.New(32),
		Status:      0,
	}
	config := Config{
		LogLevel: "info",
		Debug:    true,
		Output:   "console",
		File:     "../logs/",
		MaxAge:   7,
		AppName:  "wtf",
	}
	InitLogger(&config)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "accessLog", accessLog)
	basicLog, _ := ctx.Value(`accessLog`).(BasicLog)
	fmt.Printf("basicLog1111: %#v", basicLog)
	Warn(ctx, "debug")
	Info(ctx, accessLog)
}
