package metrics

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"time"
)

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
