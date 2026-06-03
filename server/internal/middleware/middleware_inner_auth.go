// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"encoding/base64"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strings"
)

// InnerBasicAuthentication API鉴权中间件
func InnerBasicAuthentication(r *ghttp.Request) {
	var (
		authHeader = r.GetHeader("Authorization")
		username   = "inner"
		password   = "inner@888!"
	)

	if authHeader == "" {
		r.Response.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	credentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	pair := strings.SplitN(string(credentials), ":", 2)
	if len(pair) != 2 || pair[0] != username || pair[1] != password {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	r.Middleware.Next()
}
