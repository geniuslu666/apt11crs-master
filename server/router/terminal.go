package router

import (
	"APT/internal/controller/app"
	"APT/internal/middleware"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Terminal 路由
func Terminal(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/terminal", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.TerminalIsAuth)

		group.Bind(
			app.NewTerminal().TerminalLogin,
		)
		group.Bind(

			app.NewTerminal().VerifyLog,

			app.NewTerminal().VerifyLogView,

			app.NewTerminal().CodeView,

			app.NewTerminal().CodeVerify,
		).Middleware(middleware.TerminalAuth).Middleware()
	})
}
