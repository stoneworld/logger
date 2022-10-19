package logger

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var _ NormalAccessLog = &AccessLog{}

type AccessLog struct {
	LogId     string
	BeginTime time.Time
	EndTime   time.Time

	Method      string
	Host        string
	RequestURI  string
	ContentType string
	UserAgent   string
	Post        interface{}
	Status      int

	NoticeInfo *hashmap.HashMap
	NoticeTime *hashmap.HashMap
}

func (l *AccessLog) AlarmHeader() string {
	return fmt.Sprintf("Host:%s\nRequst:%s\nTraceId%s\n", l.Host, l.RequestURI, l.GetTraceId())
}

func (l *AccessLog) BasicFields() []zap.Field {
	return []zap.Field{
		zap.String(`host`, l.Host),
		zap.String(`request`, l.RequestURI),
		zap.String(`TraceId`, l.GetTraceId()),
	}
}
func (l *AccessLog) AccessFields() []zap.Field {
	return []zap.Field{
		zap.String(`host`, l.Host), // TODO
		zap.String(`request`, l.RequestURI),
		zap.String(`method`, l.Method),
		zap.String(`contentType`, l.ContentType),
		zap.Int(`status`, l.Status),
		zap.Any(`post`, l.Post),
		zap.String(`TraceId`, l.GetTraceId()),

		zap.String(`startDate`, l.BeginTime.Format("2006-01-02T15:04:05.999Z07:00")),
		zap.String(`endDate`, l.EndTime.Format("2006-01-02T15:04:05.999Z07:00")),
		zap.Int32(`execTime`, int32(l.EndTime.Sub(l.BeginTime)/time.Millisecond)),

		zap.Any(`noticeTime`, toMap(l.NoticeTime)),
		zap.Any(`noticeInfo`, toMap(l.NoticeInfo)),
	}
}

func (l *AccessLog) GetStackByError(error) []byte {
	return nil
}

func (l *AccessLog) GetTraceId() string {
	return l.LogId
}

func (l *AccessLog) LogAddNotice(key string, v interface{}) {
	l.NoticeInfo.Set(key, v)
}

func (l *AccessLog) LogAddNoticeTime(key string, v interface{}) {
	l.NoticeTime.Set(key, v)
}

func (l *AccessLog) TimeCost() func(funcName string) {
	start := time.Now()
	return func(funcName string) {
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", float64(time.Since(start).Microseconds())/1000), 64)
		l.LogAddNoticeTime(funcName, fmt.Sprintf("%.4f", value))
	}
}

func toMap(hashMap *hashmap.HashMap) (ret map[string]interface{}) {
	if hashMap == nil {
		return
	}
	var mapLen int
	mapLen = hashMap.Len()
	if mapLen > 0 {
		ret = make(map[string]interface{}, mapLen)
		for kv := range hashMap.Iter() {
			ret[kv.Key.(string)] = kv.Value
		}
	}
	return
}
