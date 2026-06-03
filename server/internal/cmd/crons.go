package cmd

import (
	"APT/internal/consts"
	"APT/internal/crons"
	"APT/internal/dao"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	CronFeeds = &gcmd.Command{
		Name:  "CronListenFeeds",
		Usage: "start",
		Brief: "自动接收私域订单变更信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return crons.AhPublicApiFeeds(ctx)
		},
	}

	CronAvailabilities = &gcmd.Command{
		Name:  "CronSyncAvailabilities",
		Usage: "start",
		Brief: "自动接收可售房变更信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				PmsProperty []entity.PmsProperty
				PmsRoomType []entity.PmsRoomType
				puidNum     int64
				puidName    string
				tuidNum     int64
				tuidName    string
				runner      string
			)
			puidString := parser.GetOpt("puid").String()
			tuidString := parser.GetOpt("tuid").String()
			IsOptions := parser.GetOpt("op").String()
			if IsOptions == "Y" {
				if g.IsEmpty(puidString) && g.IsEmpty(tuidString) {
					// 查询物业
					if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Scan(&PmsProperty); err != nil && !errors.Is(err, sql.ErrNoRows) {
						g.Log().Error(ctx, err)
					} else {
						for i := range PmsProperty {
							fmt.Printf("序号 【%d】 --  物业ID【%s】  -- 物业名 【%s】\n", i+1, PmsProperty[i].Uid, PmsProperty[i].Name)
						}
						fmt.Printf("请选择物业序号\n")
						fmt.Scan(&puidNum)
						if puidNum != 0 {
							puidString = PmsProperty[puidNum-1].Uid
							puidName = PmsProperty[puidNum-1].Name
							fmt.Printf("选择的物业Uid是 %s\n", puidString)
							fmt.Printf("选择的物业是 %s\n", puidName)
						}
					}
					// 查询房型
					if err = dao.PmsRoomType.Ctx(ctx).Where("puid", puidString).OmitEmptyWhere().Hook(hook.PmsFindLanguageValueHook).Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
						g.Log().Error(ctx, err)
					} else {
						for i := range PmsRoomType {
							fmt.Printf("序号 【%d】 --  房型ID【%s】  -- 房型名 【%s】-- 价格PLAN_ID 【%s】\n", i+1, PmsRoomType[i].Uid, PmsRoomType[i].Name, PmsRoomType[i].RatePlanId)
						}
						fmt.Scan(&tuidNum)
						if tuidNum != 0 {
							tuidString = PmsRoomType[tuidNum-1].Uid
							tuidName = PmsRoomType[tuidNum-1].Name
						}
					}
				}
				fmt.Printf("puid:%s  name:%s\n", puidString, puidName)
				fmt.Printf("tuid:%s  name:%s\n", tuidString, tuidName)
				fmt.Println("是否执行同步库存操作   Y/N default Y")
				fmt.Scan(&runner)
			}
			if strings.ToUpper(runner) == "N" {
				fmt.Println("执行终止")
				return
			}
			if err = crons.AhPublicApiAvailabilities(ctx, puidString, tuidString); err != nil {
				g.Log().Error(ctx, err)
				return
			}
			return
		},
	}

	CronRoomTypeAndRates = &gcmd.Command{
		Name:  "CronSyncRoomTypeAndRates",
		Usage: "start",
		Brief: "自动接收房型和房态信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return crons.AhPublicApiRoomTypeAndRates(ctx)
		},
	}

	CronProperties = &gcmd.Command{
		Name:  "CronSyncProperties",
		Usage: "start",
		Brief: "同步房源信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			uidString := parser.GetOpt("uid").String()
			return crons.AhPublicApiProperties(ctx, uidString)
		},
	}
	CronFoodSettlement = &gcmd.Command{
		Name:  "CronFoodSettlement",
		Usage: "start",
		Brief: "餐饮订单结算",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/SETTLEMENT"))
			if err = service.FoodSettlementOrder().DailySettlement(ctx); err != nil {
				return
			}
			if err = service.FoodSettlementOrder().WeekSettlement(ctx); err != nil {
				return
			}
			if err = service.FoodSettlementOrder().MonthSettlement(ctx); err != nil {
				return
			}
			return
		},
	}
	CronCarSettlement = &gcmd.Command{
		Name:  "CronCarSettlement",
		Usage: "start",
		Brief: "接送机订单结算",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/SETTLEMENT"))
			if err = service.CarSettlementOrder().DailySettlement(ctx); err != nil {
				return
			}
			if err = service.CarSettlementOrder().WeekSettlement(ctx); err != nil {
				return
			}
			if err = service.CarSettlementOrder().MonthSettlement(ctx); err != nil {
				return
			}
			return
		},
	}
	CronSpaSettlement = &gcmd.Command{
		Name:  "CronSpaSettlement",
		Usage: "start",
		Brief: "按摩订单结算",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/SETTLEMENT"))
			if err = service.SpaSettlementOrder().DailySettlement(ctx); err != nil {
				return
			}
			if err = service.SpaSettlementOrder().WeekSettlement(ctx); err != nil {
				return
			}
			if err = service.SpaSettlementOrder().MonthSettlement(ctx); err != nil {
				return
			}
			return
		},
	}
	CronHotelCheckOut = &gcmd.Command{
		Name:  "CronHotelCheckOut",
		Usage: "start",
		Brief: "酒店订单手动退房",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				HotelSettingConfig  *model.HotelSettingConfig
				Logger              = g.Log().Path("logs/CRON/CronHotelCheckOut")
				expiredReservations []*entity.PmsAppReservation
				processedCount      int
				OrderSnArr          []string
				fxOrderSnArr        []string
			)
			g.DB().SetLogger(Logger)
			Logger.Info(ctx, "开始执行酒店订单自动退房任务")

			if HotelSettingConfig, err = service.BasicsConfig().GetHotelSettingConfig(ctx); err != nil {
				return
			}
			checkoutDeadline := gtime.Now().AddDate(0, 0, -HotelSettingConfig.HotelOrderCheckoutDay)

			if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				// 1.查找满足条件的酒店入住订单
				Logger.Info(ctx, "查找需要自动退房的订单，退房时间早于：", checkoutDeadline.Format("Y-m-d H:i:s"))

				mod := dao.PmsAppReservation.Ctx(ctx).TX(tx).
					Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
					Where(dao.PmsAppReservation.Columns().Status, "confirmed").
					Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_in").
					Where(dao.PmsAppReservation.Columns().Source, "APP").
					WhereLT(dao.PmsAppReservation.Columns().CheckoutDate, checkoutDeadline)

				if HotelSettingConfig.HandleCheckoutOrderNum > 0 {
					mod = mod.Limit(HotelSettingConfig.HandleCheckoutOrderNum)
				}

				if err = mod.
					Scan(&expiredReservations); err != nil && !errors.Is(err, sql.ErrNoRows) {
					Logger.Error(ctx, "查询过期订单失败：", err)
					return
				}

				if len(expiredReservations) == 0 {
					Logger.Info(ctx, "没有需要自动退房的订单")
					return
				}

				Logger.Infof(ctx, "找到 %d 个需要自动退房的订单", len(expiredReservations))
				Logger.Info(ctx, gjson.New(expiredReservations))

				// 2.批量更新订单状态为已退房
				for _, reservation := range expiredReservations {
					if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
						WherePri(reservation.Id).
						Update(g.Map{
							dao.PmsAppReservation.Columns().CheckinStatus: "checked_out",
							dao.PmsAppReservation.Columns().UpdatedAt:     gtime.Now(),
						}); err != nil {
						Logger.Errorf(ctx, "更新订单 %s 状态失败：%v", reservation.OrderSn, err)
						return
					}

					Logger.Infof(ctx, "订单 %s 已自动退房，原退房时间：%s",
						reservation.OrderSn,
						reservation.CheckoutDate.Format("Y-m-d H:i:s"))
					processedCount++

					// 写入orderSnArr数组
					found := false
					for _, str := range OrderSnArr { // 遍历切片中的每个元素进行比较
						if str == reservation.OrderSn { // 如果找到匹配的字符串，设置found为true并退出循环
							found = true
							break
						}
					}
					if !found { // 如果未找到匹配的字符串，则追加到切片中
						OrderSnArr = append(OrderSnArr, reservation.OrderSn)
					}

					if reservation.IsFx == "Y" {
						fxfound := false
						for _, str := range fxOrderSnArr { // 遍历切片中的每个元素进行比较
							if str == reservation.OrderSn { // 如果找到匹配的字符串，设置found为true并退出循环
								fxfound = true
								break
							}
						}
						if !fxfound { // 如果未找到匹配的字符串，则追加到切片中
							fxOrderSnArr = append(fxOrderSnArr, reservation.OrderSn)
						}
					}

				}

				return
			}); err != nil {
				Logger.Error(ctx, "事务执行失败：", err)
				return
			}

			// 3.投入返佣队列处理
			if !g.IsEmpty(OrderSnArr) {
				Logger.Infof(ctx, "开始将 %d 个订单投入返佣队列", len(OrderSnArr))

				Logger.Infof(ctx, "找到 %d 个分销订单需要投入返佣队列", len(fxOrderSnArr))

				for _, OrderSn := range OrderSnArr {

					if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
						ExchangeName: consts.RabbitMQExchangeName,
						QueueName:    consts.RabbitMQQueueNameRebate,
						DataByte:     gvar.New(OrderSn).Bytes(),
						Header:       nil,
					}); err != nil {
						Logger.Errorf(ctx, "订单 %s 投入返佣队列失败：%v", OrderSn, err)
						// 不返回错误，继续处理其他订单
						continue
					}

					Logger.Infof(ctx, "订单 %s 已投入返佣队列", OrderSn)

					if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
						ExchangeName: consts.RabbitMQExchangeName,
						QueueName:    consts.RabbitMQQueueNameExp,
						DataByte:     gvar.New(OrderSn).Bytes(),
						Header:       nil,
					}); err != nil {
						Logger.Errorf(ctx, "订单 %s 投入经验队列失败：%v", OrderSn, err)
						// 不返回错误，继续处理其他订单
						continue
					}
					Logger.Infof(ctx, "订单 %s 已投入经验队列", OrderSn)

				}

				for _, fxOrderSn := range fxOrderSnArr {
					_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
						ExchangeName: consts.RabbitMQExchangeName,
						QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
						DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
							OrderNo:      fxOrderSn,
							ChangeStatus: "COMPLETE",
						}).MustToJson(),
						Header: nil,
					})
				}
			}

			Logger.Infof(ctx, "酒店订单自动退房任务完成，共处理 %d 个订单", len(OrderSnArr))
			return
		},
	}
	CronCouponExp = &gcmd.Command{
		Name:  "CronCouponExp",
		Usage: "start",
		Brief: "优惠券过期",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/CronCouponExp"))
			var (
				PmsCoupons []*entity.PmsCoupon
			)
			if err = dao.PmsCoupon.Ctx(ctx).
				Where(dao.PmsCoupon.Columns().State, 1).Scan(&PmsCoupons); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}

			if len(PmsCoupons) < 0 {
				err = gerror.New("没有优惠券过期")
				g.Log().Info(ctx, "没有优惠券过期")
				return
			}

			for _, coupon := range PmsCoupons {
				if !g.IsEmpty(coupon.EndTime) && gtime.Now().After(coupon.EndTime) {
					if _, err = dao.PmsCoupon.Ctx(ctx).WherePri(coupon.Id).Update(g.Map{
						dao.PmsCoupon.Columns().State: 3,
					}); err != nil {
						return
					}
					g.Log().Infof(ctx, "优惠券ID：%d 优惠券过期", coupon.Id)
				}
			}
			g.Log().Info(ctx, "优惠券过期处理完成")
			return
		},
	}
	CronThCouponExp = &gcmd.Command{
		Name:  "CronThCouponExp",
		Usage: "start",
		Brief: "礼品券过期",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/CronThCouponExp"))
			var (
				ThMemberCoupons []*entity.ThMemberCoupon
			)
			if err = dao.ThMemberCoupon.Ctx(ctx).
				Where(dao.ThMemberCoupon.Columns().State, 2).Scan(&ThMemberCoupons); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}

			if len(ThMemberCoupons) < 0 {
				err = gerror.New("没有礼品券过期")
				g.Log().Info(ctx, "没有礼品券过期")
				return
			}

			for _, coupon := range ThMemberCoupons {
				if !g.IsEmpty(coupon.EndTime) && gtime.Now().After(coupon.EndTime) {
					if _, err = dao.ThMemberCoupon.Ctx(ctx).WherePri(coupon.Id).Update(g.Map{
						dao.ThMemberCoupon.Columns().State: 4,
					}); err != nil {
						return
					}
					g.Log().Infof(ctx, "礼品券ID：%d 礼品券过期", coupon.Id)
				}
			}
			g.Log().Info(ctx, "礼品券过期处理完成")
			return
		},
	}
	CronEmployeeActivityExp = &gcmd.Command{
		Name:  "CronEmployeeActivityExp",
		Usage: "start",
		Brief: "员工活动过期",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.DB().SetLogger(g.Log().Path("logs/CRON/CronEmployeeActivityExp"))
			var (
				EmployeeActivities []*entity.EmployeeActivity
			)
			if err = dao.EmployeeActivity.Ctx(ctx).
				Where(dao.EmployeeActivity.Columns().Status, 2).Scan(&EmployeeActivities); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}

			if len(EmployeeActivities) < 0 {
				err = gerror.New("没有员工活动过期")
				g.Log().Info(ctx, "没有员工活动过期")
				return
			}

			for _, activity := range EmployeeActivities {
				if !g.IsEmpty(activity.EndTime) && gtime.Now().After(activity.EndTime) {
					if _, err = dao.EmployeeActivity.Ctx(ctx).WherePri(activity.Id).Update(g.Map{
						dao.EmployeeActivity.Columns().Status: 3,
					}); err != nil {
						return
					}
					g.Log().Infof(ctx, "员工活动ID：%d 员工活动过期", activity.Id)
				}
			}
			g.Log().Info(ctx, "员工活动过期处理完成")
			return
		},
	}

	CronToretaSync = &gcmd.Command{
		Name:  "CronToretaSync",
		Usage: "start",
		Brief: "TORETA数据同步任务，定期检查通知日志中的失败记录并同步最新数据",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				Logger = g.Log().Path("logs/CRON/CronToretaSync")
			)

			if err = service.FoodOrder().SyncToretaNotifications(ctx, Logger); err != nil {
				return
			}
			return
		},
	}

	CronToretaTokenRefresh = &gcmd.Command{
		Name:  "CronToretaTokenRefresh",
		Usage: "start",
		Brief: "TORETA AccessToken刷新任务，定期刷新TORETA API的访问令牌",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				Logger = g.Log().Path("logs/CRON/CronToretaTokenRefresh")
			)
			g.DB().SetLogger(Logger)

			if err = service.BasicsConfig().RefreshToretaAccessToken(ctx); err != nil {
				Logger.Error(ctx, "TORETA AccessToken刷新失败:", err)
				return
			}
			Logger.Info(ctx, "TORETA AccessToken刷新成功")
			return
		},
	}
)
