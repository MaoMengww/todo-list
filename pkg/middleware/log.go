package middleware

import (
	"context"
	"time"
	"todo-list/pkg/logger"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var ServerLogMiddleware endpoint.Middleware = func(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		logger.InitLogger()
		start := time.Now()
		//获取rpc基本信息(例如调用的方法)
		ri := rpcinfo.GetRPCInfo(ctx)
		method := "unknow"
		if ri != nil && ri.To() != nil {
			method = ri.To().Method()
		}

		logId, ok := metainfo.GetPersistentValue(ctx, "log_id")
		if !ok {
			logId = "none"
		}

		logger.Infof("[%v] RPC Request | Method : %s | Req : %v", logId, method, req)

		//执行下一个中间件或最终的Handler
		err = next(ctx, req , resp)

		cost := time.Since(start)
		if err != nil {
			logger.Errorf("[%v] RPC Error | Method : %s | Cost : %v | Err : %v", logId, method, cost, err)
		} else {
			logger.Infof("[%v] RPC Success | Method : %s | Cost : %v | Resp : %v", logId, method, cost, resp)
		}
		return err
	} 
}
// ClientLogMiddleware 客户端日志中间件
var ClientLogMiddleware endpoint.Middleware = func(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		logger.InitLogger()
		logger.Infof("Client Sending Request: %+v", req)
		err = next(ctx, req, resp)
		if err != nil {
			logger.Warnf("Client Recv Error: %v", err)
		}
		return err
	}
}