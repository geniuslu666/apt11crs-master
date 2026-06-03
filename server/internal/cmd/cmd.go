package cmd

import (
	"APT/internal/gateway"
	"APT/internal/library/casbin"
	"APT/internal/middleware"
	"APT/internal/service"
	"APT/router"
	"APT/utility/opentelemetry"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Description: `默认启动Http服务`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return Http.Func(ctx, parser)
		},
	}

	Http = &gcmd.Command{
		Name:  "Http",
		Usage: "start",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化OpenTelemetry
			shutdown := opentelemetry.InitOpenTelemetry()
			defer shutdown()
			// 修改sql日志输出文件
			g.DB().SetLogger(g.Log())
			// 启动服务监控
			service.BasicsAdminMonitor().StartMonitor(ctx)

			// 初始化功能库配置
			service.BasicsConfig().InitConfig(ctx)

			// 加载超管数据
			service.BasicsAdminMember().LoadSuperAdmin(ctx)

			// 初始化casbin权限
			casbin.InitEnforcer(ctx)

			s := g.Server("APT11_SERVER")

			// 注册全局中间件
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				middleware.Ctx,
				middleware.CORS,
				middleware.Blacklist,
				middleware.HandleResponse,
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册后台路由
				router.Admin(ctx, group)
				// 注册Api路由
				router.Notify(ctx, group)
				// 注册App路由
				router.App(ctx, group)
				// 注册websocket路由
				router.WebSocket(ctx, group)
				// 注册Terminal路由
				router.Terminal(ctx, group)
				// 注册TravelStaff路由
				router.TravelStaff(ctx, group)
				// 注册分销模块
				router.Fx(ctx, group)
				// 注册H5一日游员工端
				router.H5DayTrip(ctx, group)
			})
			s.EnableAdmin()
			s.SetGraceful(true)
			s.Run()
			return nil
		},
	}

	HttpApi = &gcmd.Command{
		Name:  "HttpApi",
		Usage: "start",
		Brief: "start http api server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化OpenTelemetry
			shutdown := opentelemetry.InitOpenTelemetry()
			defer shutdown()
			// 修改sql日志输出文件
			g.DB().SetLogger(g.Log())
			// 初始化功能库配置
			service.BasicsConfig().InitConfig(ctx)
			s := g.Server("APT11_INNER_SERVER")
			s.Use(middleware.InnerBasicAuthentication)
			// 注册全局中间件
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				middleware.Ctx,
				middleware.CORS,
				middleware.Blacklist,
				middleware.HandleResponse,
			}...)

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册后台路由
				router.Api(ctx, group)
			})
			s.EnableAdmin()
			s.SetGraceful(true)
			s.Run()
			return nil
		},
	}

	HttpRestart = &gcmd.Command{
		Name:  "reload",
		Usage: "reload",
		Brief: "平滑重启Http服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				response   *gclient.Response
				serverPort = g.Cfg().MustGet(ctx, "server.address").String()[1:]
			)
			if response, err = g.Client().Get(ctx, "http://127.0.0.1:"+serverPort+"/debug/admin/restart"); err != nil {
				g.Log().Error(ctx, err)
				return
			}
			g.Dump(response.ReadAllString())
			return
		},
	}

	HttpShutdown = &gcmd.Command{
		Name:  "shutdown",
		Usage: "shutdown",
		Brief: "平滑关停Http服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				response   *gclient.Response
				serverPort = g.Cfg().MustGet(ctx, "server.address").String()[1:]
			)
			if response, err = g.Client().Get(ctx, "http://127.0.0.1:"+serverPort+"/debug/admin/shutdown"); err != nil {
				g.Log().Error(ctx, err)
				return
			}
			g.Dump(response.ReadAllString())
			return
		},
	}

	AppGateway = &gcmd.Command{
		Name:  "AppGateway",
		Usage: "start",
		Brief: "APP加密网关",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gateway.AppGateway()
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(
		Http,
		HttpApi,
		HttpRestart,
		HttpShutdown,
		AppGateway,
		MQ,
		MQOrderStay,
		MQExp,
		MQOrderExpire,
		MQPlaceOrder,
		MQAvailabilities,
		MQRebate,
		MQThCouponEffect,
		MQOrderAward,
		MQOrderRemind,
		MQOrderExport,
		MQEmployeeActivityEffect,
		MQSystemMessage,
		CronAvailabilities,
		CronProperties,
		CronFeeds,
		CronRoomTypeAndRates,
		CronFoodSettlement,
		CronCarSettlement,
		CronSpaSettlement,
		CronCouponExp,
		CronThCouponExp,
		CronEmployeeActivityExp,
		CronToretaSync,
		CronToretaTokenRefresh,
		ClearCabinetOrders,
		CronHotelCheckOut,
		MQFxChangeOrderPush,
	); err != nil {
		panic(err)
	}
}
