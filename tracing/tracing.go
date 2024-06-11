package tracing

import (
	"context"
	"errors"
	"net/http"
	"time"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	sharedLogging "github.com/wwi21seb-projekt/alpha-shared/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func StartMetricsServer(logger *zap.SugaredLogger, reg *prometheus.Registry) error {
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorHandling:     promhttp.ContinueOnError,
		EnableOpenMetrics: true,
	}))
	logger.Info("Metrics server started on port 2112")
	return http.ListenAndServe(":2112", nil)
}

func InitTelemetry(ctx context.Context, name, version string) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error
	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// Create a resource
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(name),
		semconv.ServiceVersionKey.String(version),
	)

	// Create the Jaeger exporter
	traceExp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(traceExp),
		tracesdk.WithResource(res),
	)
	shutdownFuncs = append(shutdownFuncs, tp.Shutdown)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Create prometheus exporter
	metricExp, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	// Create the metric provider
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExp)),
		metric.WithResource(res),
	)
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetMeterProvider(mp)

	return shutdown, nil
}

func GetClientMetrics() *grpcprom.ClientMetrics {
	return grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
}

func GetServerMetrics() *grpcprom.ServerMetrics {
	return grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
}

func NewServerOptions(srvMetrics *grpcprom.ServerMetrics, logger *zap.Logger, counter *uint64) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.UnaryServerInterceptor(sharedLogging.InterceptorLogger(logger), loggingOptions()...),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler(logger, counter))),
		),
	}
}

func NewClientOptions(clMetrics *grpcprom.ClientMetrics, logger *zap.Logger) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(1*time.Second),
			clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.UnaryClientInterceptor(sharedLogging.InterceptorLogger(logger), loggingOptions()...),
		),
	}
}

func grpcPanicRecoveryHandler(logger *zap.Logger, counter *uint64) recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		// Increment the counter
		*counter++
		logger.Info("recovered from panic", zap.Any("panic", p))
		return status.Errorf(codes.Internal, "%s", p)
	}
}

func exemplarFromContext(ctx context.Context) prometheus.Labels {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{"traceID": span.TraceID().String()}
	}
	return nil
}

func loggingOptions() []logging.Option {
	return []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithDurationField(logging.DefaultDurationToFields),
		logging.WithLevels(logging.DefaultServerCodeToLevel),
		logging.WithFieldsFromContext(sharedLogging.LogTraceId),
	}
}
