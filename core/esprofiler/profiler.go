package esprofiler

import (
	"compay/core/config"
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// var tracerExp *otlptrace.Exporter

func InitInstance(ctx context.Context) {
	if config.TelemetryCollectorHost == "" {
		return
	}
	optionsTrace := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(config.TelemetryCollectorHost),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
	}
	traceExporter, err := otlptracehttp.New(ctx, optionsTrace...)
	if err != nil {
		panic(err)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("esprofiler"),
		semconv.ServiceVersion(config.ProfilerVersion),
	)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}
