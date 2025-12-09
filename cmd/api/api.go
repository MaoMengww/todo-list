package main

import (
	"todo-list/app/gateway/router"
	grpc "todo-list/app/gateway/rpc"
	"todo-list/config"
	"todo-list/pkg/logger"
)

func main() {
	config.Init()
	logger.InitLogger()
	grpc.TodoInit()
	grpc.UserInit()
	router.Init()
}