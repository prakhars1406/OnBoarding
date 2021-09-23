package logger

import (
	"context"
	"time"
)

const contextKeyApiLogField = "api_log_field"

type ApiLogField struct {
	Path             string    `json:"path"`
	RequestMethod    string    `json:"request_method"`
	RequestTimestamp time.Time `json:"request_timestamp"`
	ResponseStatus   int       `json:"response_status"`
	TotalTime        int64     `json:"total_time"`
	MongoQueryTime   int64     `json:"mongo_query_time"`
	RedisQueryTime   []int64   `json:"redis_query_time"`
	HostName         string    `json:"host_name"`
}

func (f *ApiLogField) SetMongoQueryTime(t time.Duration) {
	f.MongoQueryTime = t.Milliseconds()
}

func (f *ApiLogField) SetRedisQueryTime(t time.Duration) {
	f.RedisQueryTime = append(f.RedisQueryTime, t.Milliseconds())
}

func ContextWithApiLogField(ctx context.Context, apiField *ApiLogField) context.Context {
	return context.WithValue(ctx, contextKeyApiLogField, apiField)
}

func ApiLogFieldFromContext(ctx context.Context) (*ApiLogField, bool) {
	if v := ctx.Value(contextKeyApiLogField); v != nil {
		if field, ok := v.(*ApiLogField); ok {
			return field, ok
		}
	}
	return nil, false
}
