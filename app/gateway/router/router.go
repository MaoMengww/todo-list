package router

import (
	"fmt"
	"todo-list/app/gateway/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	h := server.Default()
	user := h.Group("/api/v1/user")
	{
		user.POST("/register", handler.RegisterHandler)
		user.POST("/login", handler.LoginHandler)
		user.GET("/info", handler.UserInfoHandler)
	}
	fmt.Println("Gateway  is running on :8888")
	h.Spin()

}