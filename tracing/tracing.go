package tracing

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	sharedLogging "github.com/wwi21seb-projekt/alpha-shared/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// InitializeTelemetry initializes telemetry with OpenTelemetry.
func InitializeTelemetry(ctx context.Context, name, version string) (func(), error) {
	// Create a resource
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(name),
		semconv.ServiceVersionKey.String(version),
	)

	// Create the OTLP exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	traceProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(traceExporter),
		tracesdk.WithResource(res),
	)

	// Set up propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	otel.SetTracerProvider(traceProvider)

	shutdown := func() {
		if err = traceProvider.Shutdown(ctx); err != nil {
			fmt.Print("failed to shutdown telemetry", zap.Error(err))
		}
	}

	return shutdown, nil
}

// InitializeMetrics initializes OpenTelemetry metrics.
func InitializeMetrics(ctx context.Context, name, version string) (func(), error) {
	// Create a resource
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(name),
		semconv.ServiceVersionKey.String(version),
	)

	// Create the OTLP exporter
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(3*time.Second))),
		metric.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)

	shutdown := func() {
		if err = meterProvider.Shutdown(ctx); err != nil {
			fmt.Print("failed to shutdown telemetry", zap.Error(err))
		}
	}

	return shutdown, nil
}

// NewServerOptions returns a list of gRPC server options with tracing, logging, and recovery interceptors.
func NewServerOptions(logger *zap.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()), // Intercepts gRPC calls for telemetry
		grpc.ChainUnaryInterceptor( // Chain multiple unary interceptors
			logging.UnaryServerInterceptor(sharedLogging.InterceptorLogger(logger), loggingOptions()...),    // Logging interceptor
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler(logger))), // Recovery interceptor
		),
	}
}

// NewClientOptions returns a list of gRPC client options with tracing, logging, and timeout interceptors.
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

// grpcPanicRecoveryHandler handles panics and logs the incident.
func grpcPanicRecoveryHandler(logger *zap.Logger) recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		logger.Info("Recovered from panic", zap.Any("panic", p))
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
