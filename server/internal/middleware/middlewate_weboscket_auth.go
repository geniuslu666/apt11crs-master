// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// WebSocketAuth websocket鉴权中间件
func WebSocketAuth(r *ghttp.Request) {

	// 将用户信息传递到上下文中
	if err := DeliverUserContext(r); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, err.Error()))
		return
	}

	r.Middleware.Next()
}
