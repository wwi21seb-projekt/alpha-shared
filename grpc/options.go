package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	sharedLogging "github.com/wwi21seb-projekt/alpha-shared/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcPanicRecoveryHandler handles panics and logs the incident with a stack trace.
func grpcPanicRecoveryHandler(logger *zap.Logger) recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		logger.Error("Recovered from panic", zap.Any("panic", p), zap.Stack("stack"))
		return status.Errorf(codes.Internal, "Recovered from panic: %v", p)
	}
}

// loggingOptions returns logging interceptor options.
func loggingOptions() []logging.Option {
	return []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithDurationField(logging.DefaultDurationToFields),
		logging.WithLevels(logging.DefaultServerCodeToLevel),
		logging.WithFieldsFromContext(sharedLogging.LogTraceId),
	}
}
