// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"APT/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Blacklist IP黑名单限制中间件
func Blacklist(r *ghttp.Request) {
	if err := service.BasicsBlacklist().VerifyRequest(r); err != nil {
		r.SetError(gerror.NewCode(gerror.Code(err), err.Error()))
	}
	r.Middleware.Next()
}
