package router

import (
	"fmt"
	"todo-list/app/gateway/handler"
	"todo-list/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	h := server.Default()
	h.Use(middleware.Sentinel())
	user := h.Group("/api/v1/user")
	{
		user.POST("/register", handler.RegisterHandler)
		user.POST("/login", handler.LoginHandler)
		user.GET("/info", handler.UserInfoHandler)
	}
	todo := h.Group("/api/v1/todo", middleware.JWTAuthMiddleware())
	{
		todo.POST("/create", handler.AddTodoHandler)
		todo.GET("/list", handler.ListTodosHandler)
		todo.PUT("/update", handler.UpdateTodoHandler)
		todo.DELETE("/delete", handler.DeleteTodoHandler)
		todo.GET("/get", handler.GetTodoHandler)
	}
	fmt.Println("Gateway  is running on :8888")
	h.Spin()

}