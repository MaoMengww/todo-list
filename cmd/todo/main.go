package main

import (
	"context"
	"log"
	"net"
	"todo-list/app/todo"
	"todo-list/config"
	"todo-list/pkg/common"
	"todo-list/pkg/middleware"

	"todo-list/kitex_gen/todo/todoservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)
func main() {
	config.Init()

	shutdown := common.InitTracing("todo")
	defer shutdown(context.Background())


	todoService := todo.TodoInit()
	r, err := etcd.NewEtcdRegistry(viper.GetStringSlice("etcd.endpoints"))
	if err != nil {
		log.Fatalf("服务注册失败: %v", err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", viper.GetString("server.todo.address"))
	svr := todoservice.NewServer(
		todoService,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "todo",
		}),
		server.WithMiddleware(middleware.ServerLogMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
	)
	err = svr.Run()
	

}


