package main

import (
	"todo-list/app/gateway/router"
	grpc "todo-list/app/gateway/rpc"
	"todo-list/config"
)

func main() {
	config.Init()
	grpc.TodoInit()
	grpc.UserInit()
	router.Init()
}