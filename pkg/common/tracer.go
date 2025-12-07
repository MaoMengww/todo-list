package common

import (
	"context"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) func(context.Context) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)

	return  func(ctx context.Context) {
		p.Shutdown(ctx)
	}
}