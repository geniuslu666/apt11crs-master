package middleware

import (
	"APT/internal/library/contexts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func ServerLog(greq *ghttp.Request) {
	greq.Middleware.Next()
	contexts.SetDataMap(greq.Context(), g.Map{
		"request.takeUpTime": gtime.Now().Sub(gtime.New(greq.EnterTime)).Milliseconds(),
		// ...
	})
	// 记录日志
	AnalysisLog(greq.Context())
}
