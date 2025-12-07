package handler

import (
	"context"
	"todo-list/app/gateway/rpc"
	"todo-list/kitex_gen/user"
	"todo-list/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

 func RegisterHandler (c context.Context, ctx *app.RequestContext) {
	var req user.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.UserClient.Register(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, resp)
 }

 func LoginHandler (c context.Context, ctx *app.RequestContext) {
	var req user.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}

	resp, err := rpc.UserClient.Login(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}
	token, err := utils.GenerateJWT(resp.Info.UserId)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Token generation failed",
		})
		return
	}
	

	ctx.JSON(200, map[string]interface{}{
		"code":  200,
		"token": token,
	})
 }

 func UserInfoHandler (c context.Context, ctx *app.RequestContext) {
	var req user.GetUserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(400, map[string]string{
			"error": "Invalid request",
		})
		return
	}
	resp, err := rpc.UserClient.GetUser(c, &req)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code":  "200",
		"userId": resp.Info.UserId,
	})
 }
