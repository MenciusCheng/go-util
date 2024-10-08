package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

var (
	logger      *zap.Logger
	atomicLevel zap.AtomicLevel
)

// init 函数在包导入时自动执行
func init() {
	// 创建 AtomicLevel 实例，默认为 Info 级别
	atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.Config{
		Encoding:         "json",      // JSON 格式
		Level:            atomicLevel, // 使用 AtomicLevel
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "ts", // 使用 ts 作为时间字段
			LevelKey:   "level",
			MessageKey: "msg",
			//CallerKey:     "caller", // 打印调用者文件与行号
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写输出日志级别
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			//EncodeCaller:  zapcore.FullCallerEncoder, // 打印完整的调用函数路径
		},
	}

	var err error
	logger, err = config.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

// Debug 打印 debug 级别的日志
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Debug(msg, addTraceInfo(ctx, fields...)...)
}

// Info 打印 info 级别的日志
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, addTraceInfo(ctx, fields...)...)
}

// Warn 打印 warn 级别的日志
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, addTraceInfo(ctx, fields...)...)
}

// Error 打印 error 级别的日志
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, addTraceInfo(ctx, fields...)...)
}

// addTraceInfo 添加traceId、代码行、时间等信息
func addTraceInfo(ctx context.Context, fields ...zap.Field) []zap.Field {

	_, file, line, _ := runtime.Caller(2)
	fields = append(fields, zap.String("defineCaller", fmt.Sprintf("%s:%d", file, line)))

	if ctx != nil {
		traceID := ctx.Value("traceId")
		if traceID != nil {
			fields = append(fields, zap.Any("traceId", traceID))
		}
	}

	return fields
}

// SetLogLevel 设置日志级别
func SetLogLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

// Close 关闭日志
func Close() {
	_ = logger.Sync()
}
