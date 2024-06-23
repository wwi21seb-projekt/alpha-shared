package logging

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func InitializeLogger(serviceName string) (*zap.SugaredLogger, func()) {
	var config zap.Config
	if os.Getenv("ENVIRONMENT") == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Build the logger from the configuration
	logger, err := config.Build()
	if err != nil {
		// Build a simple logger if the configuration fails
		logger = zap.NewExample()
		logger.Error("Failed to build logger from configuration", zap.Error(err))
	}

	// Create a sugared logger and add service context
	sugaredLogger := logger.Sugar().With(zap.String("service", serviceName))

	// Cleanup function to flush the logger buffer
	cleanup := func() {
		if err = sugaredLogger.Sync(); err != nil {
			sugaredLogger.Errorf("Failed to sync logger: %v", err)
		}
	}

	return sugaredLogger, cleanup
}

// InterceptorLogger adapts zap logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func LogTraceId(ctx context.Context) logging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return logging.Fields{"traceID", span.TraceID().String()}
	}
	return nil
}
