package main

import (
	"context"
	"log"
	"net"
	"todo-list/app/user"
	"todo-list/config"
	"todo-list/pkg/common"
	"todo-list/pkg/middleware"

	"todo-list/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)
func main() {
	config.Init()

	shutdown := common.InitTracing("user")
	defer shutdown(context.Background())


	userService := user.UserInit()
	r, err := etcd.NewEtcdRegistry(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		log.Fatalf("服务注册失败: %v", err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", viper.GetString("server.user.address"))
	svr := userservice.NewServer(
		userService,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user",
		}),
		server.WithMiddleware(middleware.ServerLogMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
	)
	err = svr.Run()
	

}


