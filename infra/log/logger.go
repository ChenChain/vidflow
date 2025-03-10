package log

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	TraceIDKey = "trace_id"
)

// Logger 是一个支持分布式追踪的日志类
type Logger struct {
	zapLogger *zap.Logger
}

// 全局 Logger 实例（可以作为单例调用）
var globalLogger *Logger

// 初始化全局 Logger
func InitGlobalLogger(level zapcore.Level, enableTrace bool) {
	globalLogger = NewLogger(level, enableTrace)
}

// 返回全局 Logger 实例
func GlobalLogger() *Logger {
	return globalLogger
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger(level zapcore.Level, enableTrace bool) *Logger {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false, // 是否为开发模式
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			MessageKey:    "msg",
			NameKey:       "logger",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	// 如果启用分布式追踪，则将 Trace 信息添加为 Zap 默认上下文
	if enableTrace {
		logger = logger.With(zap.String("service_name", "vidflow"))
	}

	return &Logger{
		zapLogger: logger,
	}
}

// 获取 Trace ID 和 Span ID 从 Context
func extractTraceInfo(ctx context.Context) string {
	traceID, _ := ctx.Value(TraceIDKey).(string)
	return traceID
}

// 包裹 Zap 日志方法，将 Trace 信息添加到日志
func (l *Logger) logWithTrace(ctx context.Context, level zapcore.Level, msg string, fields ...zap.Field) {
	traceID := extractTraceInfo(ctx)

	// 添加 Trace ID 和 Span ID 信息
	logFields := append(fields, zap.String(TraceIDKey, traceID))

	// 根据日志级别记录日志
	switch level {
	case zapcore.DebugLevel:
		l.zapLogger.Debug(msg, logFields...)
	case zapcore.InfoLevel:
		l.zapLogger.Info(msg, logFields...)
	case zapcore.WarnLevel:
		l.zapLogger.Warn(msg, logFields...)
	case zapcore.ErrorLevel:
		l.zapLogger.Error(msg, logFields...)
	case zapcore.FatalLevel:
		l.zapLogger.Fatal(msg, logFields...)
	}
}

// 公共日志方法

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.logWithTrace(ctx, zapcore.DebugLevel, msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.logWithTrace(ctx, zapcore.InfoLevel, msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.logWithTrace(ctx, zapcore.WarnLevel, msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.logWithTrace(ctx, zapcore.ErrorLevel, msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.logWithTrace(ctx, zapcore.FatalLevel, msg, fields...)
}

// 示例：将 Trace 信息传递到 Context
func WithTrace(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, TraceIDKey, traceID)
	return ctx
}

// 自动生成 Trace ID (可以直接用 UUID)
func generateTraceID() string {
	return uuid.New().String() // 生成一个 UUID 作为 Trace ID
}
