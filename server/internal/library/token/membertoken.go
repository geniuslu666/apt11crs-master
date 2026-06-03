// Package token
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package token

import (
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/utility/guomi"
	"APT/utility/simple"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MemberClaims struct {
	*model.MemberIdentity
	jwt.RegisteredClaims
}

type MemberToken struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}

// MemberLogin 登录
func MemberLogin(ctx context.Context, user *model.MemberIdentity) (string, int64, error) {
	var (
		jsonData []byte
		header   []byte
		err      error
		now      *gtime.Time
		authKey  string
		tokenKey string
		bindKey  string
		duration time.Duration
		token    *Token
		claims   MemberClaims
	)
	claims = MemberClaims{
		user,
		jwt.RegisteredClaims{},
	}
	// 将结构体编码为JSON格式
	if jsonData, err = json.Marshal(claims); err != nil {
		goto ERR
	}
	if header, err = guomi.NewSM4(sm4key, sm4Iv).SM4En(jsonData); err != nil {
		goto ERR
	}
	now = gtime.Now()
	// 认证key
	authKey = GetAuthKey(string(header))
	// 登录token
	tokenKey = GetTokenKey(user.App, authKey)
	// 身份绑定
	bindKey = GetBindKey(user.App, gvar.New(user.Id).Int64())
	// 有效时长
	duration = time.Second * gconv.Duration(config.Expires)
	token = &Token{
		ExpireAt:     now.Unix() + config.Expires,
		RefreshAt:    now.Unix(),
		RefreshCount: 0,
	}
	if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
		goto ERR
	}
	if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
		goto ERR
	}
	return string(header), config.Expires, nil
ERR:
	return "", 0, err
}

// MemberLogout 注销登录
func MemberLogout(r *ghttp.Request) (err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := parseToken(ctx, header)
	if err != nil {
		g.Log().Debugf(ctx, "logout parseToken err:%+v", err)
		err = errorLogin
		return
	}

	var (
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(contexts.GetModule(ctx), authKey)
		// 身份绑定
		bindKey = GetBindKey(contexts.GetModule(ctx), claims.Id)
	)

	// 删除token
	if _, err = cache.Instance().Remove(ctx, tokenKey); err != nil {
		return
	}

	if !config.MultiLogin {
		if _, err = cache.Instance().Remove(ctx, bindKey); err != nil {
			return
		}
	}
	return
}

// MemberParseLoginUser 解析登录用户信息
func MemberParseLoginUser(r *ghttp.Request) (user *model.MemberIdentity, err error) {
	var (
		ctx    = r.Context()
		header = GetAuthorization(r)
	)

	if header == "" {
		err = errorLogin
		return
	}

	claims, err := MemberparseToken(ctx, header)
	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		err = errorLogin
		return
	}

	var (
		// 认证key
		authKey = GetAuthKey(header)
		// 登录token
		tokenKey = GetTokenKey(claims.App, authKey)
		// 身份绑定
		bindKey = GetBindKey(claims.App, gvar.New(claims.Id).Int64())
	)

	// 检查token是否存在
	tk, err := cache.Instance().Get(ctx, tokenKey)
	if err != nil {
		g.Log().Debugf(ctx, "get tokenKey err:%+v", err)
		err = errorLogin
		return
	}

	if tk.IsEmpty() {
		g.Log().Debug(ctx, "token isEmpty")
		err = errorLogin
		return
	}

	var token *Token
	if err = tk.Scan(&token); err != nil {
		g.Log().Debugf(ctx, "token scan err:%+v", err)
		err = errorLogin
		return
	}

	if token == nil {
		g.Log().Debugf(ctx, "token = nil")
		err = errorLogin
		return
	}

	now := gtime.Now()
	if token.ExpireAt < now.Unix() {
		g.Log().Debugf(ctx, "token expired.")
		err = errorLogin
		return
	}

	// 是否允许多端登录
	if !config.MultiLogin {
		origin, err := cache.Instance().Get(ctx, bindKey)
		if err != nil {
			g.Log().Debugf(ctx, "bindKey get err:%+v", err)
			err = errorLogin
			return nil, err
		}

		if origin == nil || origin.IsEmpty() {
			g.Log().Debug(ctx, "bindKey isEmpty")
			err = errorLogin
			return nil, err
		}

		if tokenKey != origin.String() {
			g.Log().Debugf(ctx, "bindKey offsite login tokenKey:%v, origin:%v", tokenKey, origin.String())
			err = errorMultiLogin
			return nil, err
		}
	}

	// 自动刷新token有效期
	refreshToken := func() {
		// 未开启自动刷新
		if !config.AutoRefresh {
			return
		}

		// 刷新次数已达上限
		if config.MaxRefreshTimes != -1 && token.RefreshCount >= config.MaxRefreshTimes {
			return
		}

		// 未达到刷新间隔
		if gtime.New(token.RefreshAt).Unix()+config.RefreshInterval > now.Unix() {
			return
		}

		// 刷新有效期
		token.ExpireAt = now.Unix() + config.Expires
		token.RefreshAt = now.Unix()
		token.RefreshCount += 1

		duration := time.Second * gconv.Duration(config.Expires)

		if err = cache.Instance().Set(ctx, tokenKey, token, duration); err != nil {
			return
		}

		if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
			return
		}
	}

	simple.SafeGo(ctx, func(ctx context.Context) {
		refreshToken()
	})

	user = claims.MemberIdentity
	return
}

// MemberparseToken 解析jwt令牌
func MemberparseToken(ctx context.Context, header string) (*MemberClaims, error) {
	var (
		token  []byte
		claims *MemberClaims
		err    error
	)

	if token, err = guomi.NewSM4(sm4key, sm4Iv).SM4De([]byte(header)); err != nil {
		goto ERR
	}

	// 将结构体编码为JSON格式
	if err = json.Unmarshal(token, &claims); err != nil {
		goto ERR
	}
	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		return nil, err
	}
	return claims, nil
ERR:
	g.Log().Debugf(ctx, "parseToken err:%+v", err)
	return nil, err
}
