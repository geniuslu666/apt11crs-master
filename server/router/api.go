// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"APT/internal/controller/inner"
	"APT/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Api(ctx context.Context, group *ghttp.RouterGroup) {
	group.Middleware(middleware.ServerLog)

	group.GET("/inner", inner.NewOrder().OrderList)

	// 注册生成路由
	//genrouter.Register(ctx, group)
}
