package router

import (
	"APT/internal/controller/app"
	"APT/internal/middleware"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// TravelStaff 路由
func TravelStaff(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/travelStaff", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.TravelStaffIsAuth)

		group.Bind(
			app.NewTravel().StaffLogin,
			app.NewTravel().StaffLogout,
			app.NewTravel().StaffConfig,
		)
		group.Bind(

			app.NewTravel().VerifyLog,
			app.NewTravel().VerifyLogView,
			app.NewTravel().CodeView,
			app.NewTravel().CodeVerify,
		).Middleware(middleware.TravelStaffAuth).Middleware()
	})
}
