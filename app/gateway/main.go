package gateway

import (
	"context"
	"todo-list/app/gateway/mw"
	"todo-list/app/gateway/rpc"
	"todo-list/pkg/common"

	"github.com/cloudwego/hertz/pkg/app/server"

	hertxtracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
)

func main() {
	shutdown := common.InitTracing("api-gateway")
	defer shutdown(context.Background())

	rpc.Init()

	tracer, config := hertxtracing.NewServerTracer()


	h := server.Default(
		server.WithHostPorts(":8080"),
		tracer,
	)
	
	h.Use(hertxtracing.ServerMiddleware(config))
	h.Use(mw.Sentinel())

	h.Spin()

	
}