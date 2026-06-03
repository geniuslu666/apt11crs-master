package admin

import (
	"APT/internal/dao"
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
)

func (c *ControllerPms) Dashboard(ctx context.Context, req *pms.DashboardReq) (res *pms.DashboardRes, err error) {
	// 今日
	todayDate := gtime.Now().Format("Y-m-d")

	res = new(pms.DashboardRes)
	if req.Type == "base" {
		//今日预抵
		if res.Details.ImportBookings, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			Where(dao.PmsAppReservation.Columns().CheckinDate, todayDate).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "before_checkin").
			OmitEmptyWhere().Count(); err != nil {
			return
		}
		// 已办理入住
		if res.Details.CheckInBookings, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			Where(dao.PmsAppReservation.Columns().CheckinDate, todayDate).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_in").
			OmitEmptyWhere().Count(); err != nil {
			return
		}
		// 今日预离
		if res.Details.ExportBookings, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			Where(dao.PmsAppReservation.Columns().CheckoutDate, todayDate).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_in").
			OmitEmptyWhere().Count(); err != nil {
			return
		}
		// 已办理退房
		if res.Details.CheckOutBookings, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			Where(dao.PmsAppReservation.Columns().CheckoutDate, todayDate).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_out").
			OmitEmptyWhere().Count(); err != nil {
			return
		}
		// 在住客房数 todo
		if res.Details.CheckedBookings, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			WhereLTE(dao.PmsAppReservation.Columns().CheckinDate, todayDate).
			WhereGTE(dao.PmsAppReservation.Columns().CheckoutDate, todayDate).
			WhereIn(dao.PmsAppReservation.Columns().CheckinStatus, g.Slice{"checked_in"}).
			OmitEmptyWhere().Count(); err != nil {
			return
		}
		// 客房总数
		// 根据物业ID找到房型ids
		roomTypeUids, _ := service.HotelService().RoomTypeGetUidsByPuid(ctx, req.Puid)
		if res.Details.AllRoom, err = dao.PmsRoomUnit.Ctx(ctx).
			WhereIn(dao.PmsRoomUnit.Columns().RtUid, roomTypeUids).
			Count(); err != nil {
			return
		}

		// 今日预定 todo
		if res.Details.ToDayStays, err = dao.PmsAppStay.Ctx(ctx).
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		// App预定 todo
		if res.Details.ToDayAppStays, err = dao.PmsAppStay.Ctx(ctx).
			Where(dao.PmsAppStay.Columns().Source, "APP").
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		return
	}
	if req.Type == "amount" {
		var (
			transactionOrm    = dao.PmsTransaction.Ctx(ctx).LeftJoinOnFields(dao.PmsAppStay.Table(), dao.PmsTransaction.Columns().OrderSn, "=", dao.PmsAppStay.Columns().OrderSn)
			refundOrm         = dao.PmsTransactionRefund.Ctx(ctx).LeftJoinOnFields(dao.PmsAppStay.Table(), dao.PmsTransactionRefund.Columns().OrderSn, "=", dao.PmsAppStay.Columns().OrderSn)
			totalBalIncome    float64
			totalCouponIncome float64
			YYConfig          *model.YYConfig
		)

		// 获取配置
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		if g.IsEmpty(YYConfig) {
			err = gerror.New("系统异常")
			return
		}
		res.AmountStat.ExchangeRate = YYConfig.ExchangeRate

		// 总收入
		if res.AmountStat.TotalIncome, err = transactionOrm.Safe().
			WhereBetween(dao.PmsTransaction.Columns().PayTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			OmitEmptyWhere().Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		// 积分支付金额
		if totalBalIncome, err = transactionOrm.Safe().
			WhereBetween(dao.PmsTransaction.Columns().PayTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Where(dao.PmsTransaction.Columns().PayType, "BAL").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			OmitEmptyWhere().Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		// 优惠券支付金额
		if totalCouponIncome, err = transactionOrm.Safe().
			WhereBetween(dao.PmsTransaction.Columns().PayTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Where(dao.PmsTransaction.Columns().PayType, "COUPON").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			OmitEmptyWhere().Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}

		// PAYCLOUD收入
		if res.AmountStat.PayCloudIncome, err = transactionOrm.Safe().
			WhereBetween(dao.PmsTransaction.Columns().PayTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Where(dao.PmsTransaction.Columns().PayChannel, "PAYCLOUD").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			OmitEmptyWhere().Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		// PAYCLOUD退款
		//if cTransactionSnArr, err = service.SysPmsTransaction().GetTransactionSn(ctx, req.StartDate, req.EndDate, "PAYCLOUD", ""); err != nil {
		//	err = gerror.Wrap(err, "获取支付流水号失败！")
		//	return
		//}

		if res.AmountStat.PayCloudRefund, err = refundOrm.Safe().
			WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			Where(dao.PmsTransactionRefund.Columns().RefundChannel, "PAYCLOUD").
			OmitEmptyWhere().Sum("hg_pms_transaction_refund.refund_amount"); err != nil {
			return
		}

		// Strip信用卡收入
		if res.AmountStat.StripIncome, err = transactionOrm.Safe().
			WhereBetween(dao.PmsTransaction.Columns().PayTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			Where(dao.PmsTransaction.Columns().PayChannel, "STRIPE").
			OmitEmptyWhere().Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}

		// Strip信用卡退款
		//if transactionSnArr, err = service.SysPmsTransaction().GetTransactionSn(ctx, req.StartDate, req.EndDate, "STRIPE", ""); err != nil {
		//	err = gerror.Wrap(err, "获取支付流水号失败！")
		//	return
		//}

		if res.AmountStat.StripRefund, err = refundOrm.Safe().
			WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			Where(dao.PmsTransactionRefund.Columns().RefundChannel, "STRIPE").
			OmitEmptyWhere().Sum("hg_pms_transaction_refund.refund_amount"); err != nil {
			return
		}

		// 积分抵扣支付
		res.AmountStat.BalIncome = totalBalIncome

		// 积分抵扣退款
		//if balTransactionSnArr, err = service.SysPmsTransaction().GetTransactionSn(ctx, req.StartDate, req.EndDate, "", "BAL"); err != nil {
		//	err = gerror.Wrap(err, "获取支付流水号失败！")
		//	return
		//}
		if res.AmountStat.BalRefund, err = refundOrm.Safe().
			WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			Where(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			OmitEmptyWhere().Sum("hg_pms_transaction_refund.refund_amount"); err != nil {
			return
		}

		// 优惠券抵扣
		res.AmountStat.CouponIncome = totalCouponIncome

		// 优惠券退款
		//if couponTransactionSnArr, err = service.SysPmsTransaction().GetTransactionSn(ctx, req.StartDate, req.EndDate, "", "COUPON"); err != nil {
		//	err = gerror.Wrap(err, "获取支付流水号失败！")
		//	return
		//}
		if res.AmountStat.CouponRefund, err = refundOrm.Safe().
			WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, req.StartDate, req.EndDate).
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, req.Puid).
			Where(dao.PmsTransactionRefund.Columns().RefundType, "COUPON").
			OmitEmptyWhere().Sum("hg_pms_transaction_refund.refund_amount"); err != nil {
			return
		}
		// 有效收入
		res.AmountStat.EffectIncome = res.AmountStat.TotalIncome - totalBalIncome - res.AmountStat.BalRefund - totalCouponIncome - res.AmountStat.CouponRefund
	}
	// 近七天入住率
	if req.Type == "echarts" {
		if g.IsEmpty(res.Details.AllRoom) {
			// 客房总数
			// 根据物业ID找到房型ids
			roomTypeUids, _ := service.HotelService().RoomTypeGetUidsByPuid(ctx, req.Puid)
			if res.Details.AllRoom, err = dao.PmsRoomUnit.Ctx(ctx).OmitEmptyWhere().
				WhereIn(dao.PmsRoomUnit.Columns().RtUid, roomTypeUids).
				Count(); err != nil {
				return
			}
		}
		minCheckintime := gtime.Now().Add(-6 * gtime.H * 24).Format("Y-m-d")
		Fields := fmt.Sprintf("checkin_date as date,ROUND(COUNT(1) / %v * 100,2) as percent", res.Details.AllRoom)
		if err = dao.PmsAppReservation.Ctx(ctx).
			WhereGTE(dao.PmsAppReservation.Columns().CheckinDate, minCheckintime).
			WhereLTE(dao.PmsAppReservation.Columns().CheckinDate, gtime.Now().Format("Y-m-d")).
			WhereIn(dao.PmsAppReservation.Columns().CheckinStatus, g.Slice{"checked_in", "checked_out"}).
			Where(dao.PmsAppReservation.Columns().Puid, req.Puid).
			Group(dao.PmsAppReservation.Columns().CheckinDate).Fields(Fields).
			Scan(&res.CheckInTrend); err != nil {
			return
		}
	}
	return
}

func (c *ControllerPms) DashboardAll(ctx context.Context, req *pms.DashboardAllReq) (res *pms.DashboardAllRes, err error) {
	res = new(pms.DashboardAllRes)
	if res.BasicDashboardModel, err = service.BasicsDashboard().GetBasics(ctx, &req.BasicDashboardInp); err != nil {
		return
	}
	if res.RankingModel, err = service.BasicsDashboard().GetRanking(ctx, &req.RankingInp); err != nil {
		return
	}
	return
}
