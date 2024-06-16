package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	sharedLogging "github.com/wwi21seb-projekt/alpha-shared/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

// NewClientOptions returns a list of grpc.DialOptions for a client
func NewClientOptions(logger *zap.Logger) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(3*time.Second),
			logging.UnaryClientInterceptor(sharedLogging.InterceptorLogger(logger), loggingOptions()...),
		),
	}
}
