package handler

import (
	"context"
	"todo-list/app/gateway/rpc"
	"todo-list/kitex_gen/todo"

	"github.com/cloudwego/hertz/pkg/app"
)



func AddTodoHandler(c context.Context, ctx *app.RequestContext) {
	var req todo.AddTodoRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.TodoClient.AddTodo(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
}

func DeleteTodoHandler(c context.Context, ctx *app.RequestContext) {
	var req todo.DeleteTodoRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.TodoClient.DeleteTodo(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
}

func UpdateTodoHandler(c context.Context, ctx *app.RequestContext) {
	var req todo.UpdateTodoRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.TodoClient.UpdateTodo(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
}

func GetTodoHandler(c context.Context, ctx *app.RequestContext) {
	var req todo.GetTodoRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.TodoClient.GetTodo(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
}	

func ListTodosHandler(c context.Context, ctx *app.RequestContext) {
	var req todo.ListTodoRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.TodoClient.ListTodo(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
}
