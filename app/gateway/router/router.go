package router

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterRouter(h *server.Hertz) {
	h.GET("/register", func(c context.Context, ctx *app.RequestContext) {
		
	})
}