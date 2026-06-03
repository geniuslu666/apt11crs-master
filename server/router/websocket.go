package router

import (
	controller "APT/internal/controller/websocket"
	"APT/internal/controller/websocket/handler/admin"
	"APT/internal/controller/websocket/handler/common"
	"APT/internal/middleware"
	"APT/internal/websocket"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// WebSocket ws路由配置
func WebSocket(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware()
		group.Bind(
			controller.NewBasics().SendToTag, // 通过http发送ws消息。方便测试没有放权限中间件，实际使用时请自行调整
		)

		// ws连接中间件
		group.Middleware(middleware.WebSocketAuth)

		// ws
		group.GET("/socket", websocket.WsPage)
	})

	// 启动websocket监听
	websocket.Start()

	// 注册消息路由
	websocket.RegisterMsg(websocket.EventHandlers{
		"ping":                  common.Site.Ping,      // 心跳
		"join":                  common.Site.Join,      // 加入组
		"quit":                  common.Site.Quit,      // 退出组
		"admin/monitor/trends":  admin.Monitor.Trends,  // 后台监控，动态数据
		"admin/monitor/runInfo": admin.Monitor.RunInfo, // 后台监控，运行信息
	})
}
