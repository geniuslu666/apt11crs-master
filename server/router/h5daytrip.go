package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// H5DayTrip H5 一日游员工端路由
func H5DayTrip(ctx context.Context, group *ghttp.RouterGroup) {
	// 将 /daytrip/* 的非 API 请求都返回 index.html（SPA 路由支持）
	group.GET("/daytrip/*path", func(r *ghttp.Request) {
		r.Response.ServeFile("resource/public/daytrip/index.html")
	})
}
