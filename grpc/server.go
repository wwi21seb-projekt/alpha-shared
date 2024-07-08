package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	sharedLogging "github.com/wwi21seb-projekt/alpha-shared/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// NewServerOptions returns a list of gRPC server options with tracing, logging, and recovery interceptors.
func NewServerOptions(logger *zap.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()), // Intercepts gRPC calls for telemetry
		grpc.MaxRecvMsgSize(1024*5),					// Set maximum message size to 5 MB
		grpc.ChainUnaryInterceptor( // Chain multiple unary interceptors
			logging.UnaryServerInterceptor(sharedLogging.InterceptorLogger(logger), loggingOptions()...),    // Logging interceptor
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler(logger))), // Recovery interceptor
		),
	}
}
