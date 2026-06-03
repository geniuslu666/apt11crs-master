// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
	"time"
)

// ApiAuth API鉴权中间件
func ApiAuth(r *ghttp.Request) {

	if !strings.HasPrefix(r.URL.Path, "/api") {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := MemberDeliverUserContext(r); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, err.Error()))
		return
	}

	r.Middleware.Next()
}

// ApiIsAuth 是否存在授权信息 存在的话解析授权信息
func ApiIsAuth(r *ghttp.Request) {
	var (
		Authorization = r.Header.Get("Authorization")
	)
	if !g.IsEmpty(Authorization) {
		_ = MemberDeliverUserContext(r)
	}
	r.Middleware.Next()
}

func ApiAuthCheckMember(r *ghttp.Request) {
	var (
		ctx        = r.Context()
		MemberInfo = contexts.GetMemberUser(ctx)
		PmsMember  *entity.PmsMember
		err        error
	)
	if err = dao.PmsMember.Ctx(ctx).
		Where(dao.PmsMember.Columns().Id, MemberInfo.Id).
		Cache(gdb.CacheOption{
			Duration: time.Second * 60 * 1,
			Name:     fmt.Sprintf("middleware_ApiAuthCheckMember_%d", MemberInfo.Id),
			Force:    false,
		}).
		Scan(&PmsMember); err != nil {
		r.SetError(gerror.NewCode(gcode.CodeSecurityReason, "用户不存在"))
		return
	}
	// TODO 判断用户是否注销  判断用户是否已禁用
	if g.IsEmpty(PmsMember) {
		r.SetError(gerror.NewCode(gcode.CodeSecurityReason, "用户已注销"))
		return
	}

	// 判断会员启用禁用状态
	if PmsMember.Status == 2 {
		r.SetError(gerror.NewCode(gcode.CodeSecurityReason, "会员已被禁用"))
		return
	}

	r.Middleware.Next()
}
