package log

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	SetLogLevel(zap.DebugLevel)
	// 模拟上下文，带 traceId
	ctx := context.WithValue(context.Background(), "traceId", "123456")
	Debug(ctx, "这是 Debug 消息", zap.String("user", "test"))
	Info(ctx, "这是 Info 消息", zap.String("user", "test"))
	Error(ctx, "这是 Error 消息", zap.Error(errors.New("is error")), zap.String("user", "test"))
}

func TestInfo(t *testing.T) {
	ctx := context.Background()
	Debug(ctx, "这是 Debug 消息", zap.String("user", "test"))
	Info(ctx, "这是 Info 消息", zap.String("user", "test"))
	Error(ctx, "这是 Error 消息", zap.Error(errors.New("is error")), zap.String("user", "test"))
}
