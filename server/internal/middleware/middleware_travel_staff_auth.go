// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// TravelStaffAuth 一日游核销端API鉴权中间件
func TravelStaffAuth(r *ghttp.Request) {

	// 将用户信息传递到上下文中
	if err := TravelStaffDeliverUserContext(r); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, err.Error()))
		return
	}

	r.Middleware.Next()
}

// TravelStaffIsAuth 是否存在授权信息 存在的话解析授权信息
func TravelStaffIsAuth(r *ghttp.Request) {
	var (
		Authorization = r.Header.Get("Authorization")
	)
	if !g.IsEmpty(Authorization) {
		_ = TravelStaffDeliverUserContext(r)
	}
	r.Middleware.Next()
}
