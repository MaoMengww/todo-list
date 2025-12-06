package rpc

import (
	"log"
	"time"
	"todo-list/kitex_gen/user/userservice"
	"todo-list/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)

var UserClient userservice.Client
func UserInit() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.address")})
	if err != nil {
		log.Fatalf("failed to create etcd resolver: %v", err)
	}

	cbSuite := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return ri.To().ServiceName() + ":" + ri.To().Method()
	})

	c, err := userservice.NewClient(
		"user",
		client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithMiddleware(middleware.ClientLogMiddleware),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithCircuitBreaker(cbSuite),
		client.WithSuite(tracing.NewClientSuite()),
	)

	if err != nil {
		log.Fatalf("init client failed: err:%v", err)
	}

	UserClient = c
}