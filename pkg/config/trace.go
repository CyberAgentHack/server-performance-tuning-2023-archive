package config

import (
	"context"
	"time"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ConfigureTraceProvider(logger *zap.Logger) (func(), error) {
	ctx := context.Background()
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("server-performance-tuning-2023"))
	endpoint := "0.0.0.0:4317"
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithDialOption(grpc.WithBlock()))

	if err != nil {
		return nil, err
	}

	idg := xray.NewIDGenerator()

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithIDGenerator(idg),
	)

	otel.SetTracerProvider(tp)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			logger.Error("failed to shutdown TracerProvider", zap.Error(err))
		}
	}

	return cleanup, nil
}
