// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"APT/internal/consts"
	"APT/internal/library/contexts"
	"APT/internal/service"
	"APT/utility/simple"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// AdminAuth 后台鉴权中间件
func AdminAuth(r *ghttp.Request) {
	// 将用户信息传递到上下文中
	if err := DeliverUserContext(r); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, err.Error()))
		return
	}

	r.Middleware.Next()
}

// AdminRole 验证路由访问权限
func AdminRole(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, consts.AppAdmin), "", 1)
	)
	if !service.BasicsAdminRole().Verify(ctx, path, r.Method) {
		g.Log().Debugf(ctx, "AdminAuth fail path:%+v, GetRoleKey:%+v, r.Method:%+v", path, contexts.GetRoleKey(ctx), r.Method)
		// 清空响应
		r.SetError(gerror.NewCode(gcode.CodeSecurityReason, "你没有访问权限！"))
		return
	}
	r.Middleware.Next()
}
