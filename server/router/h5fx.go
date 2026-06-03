// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"APT/internal/controller/app"
	"APT/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Fx(ctx context.Context, group *ghttp.RouterGroup) {
	group.Middleware(middleware.ServerLog)
	group.Group("/fx", func(group *ghttp.RouterGroup) {
		// 验证登录权限 不验证路由地址权限
		group.Bind(
			app.NewFx().Login,
		)
	})

	// 注册生成路由
	//genrouter.Register(ctx, group)
}
