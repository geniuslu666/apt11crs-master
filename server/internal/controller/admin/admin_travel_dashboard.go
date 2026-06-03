package admin

import (
	"APT/api/admin/travel"
	"APT/internal/dao"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerTravel) Dashboard(ctx context.Context, req *travel.DashboardReq) (res *travel.DashboardRes, err error) {
	res = new(travel.DashboardRes)
	today := gtime.Now().Format("Y-m-d")

	switch req.Type {
	case "base":
		// 今日预约总数
		if res.Details.TodayOrderTotalNum, err = dao.TravelOrder.Ctx(ctx).
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
		// 今日有效预约数（非取消/退款）
		if res.Details.TodayEffectiveOrderNum, err = dao.TravelOrder.Ctx(ctx).
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "CANCEL", "REFUND"}).
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
		// 今日营业额（已支付且非取消/退款）
		if res.Details.TodayOrderAmount, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().PayStatus, "HAVE_PAID").
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "CANCEL", "REFUND"}).
			Where("DATE(created_at) = ?", today).
			Sum(dao.TravelOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 累计营业额（已支付且非取消/退款）
		if res.Details.TotalOrderAmount, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().PayStatus, "HAVE_PAID").
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "CANCEL", "REFUND"}).
			Sum(dao.TravelOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 累计预约数
		if res.Details.TotalOrderNum, err = dao.TravelOrder.Ctx(ctx).Count(); err != nil {
			return
		}
		// 累计有效预约数（非取消/退款）
		if res.Details.EffectiveOrderNum, err = dao.TravelOrder.Ctx(ctx).
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "CANCEL", "REFUND"}).
			Count(); err != nil {
			return
		}
		// 今日待核销数（预约日期=今天且状态为WAIT_VERIFY）
		if res.Details.TodayWaitVerifyNum, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "WAIT_VERIFY").
			Where("DATE(book_date) = ?", today).
			Count(); err != nil {
			return
		}
		// 今日已核销数（核销时间=今天且状态为DONE）
		if res.Details.TodayVerifiedNum, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "DONE").
			Where("DATE(verify_time) = ?", today).
			Count(); err != nil {
			return
		}

	case "todayOrder":
		// 待支付订单数（今日创建）
		if res.TodayOrder.WaitPayOrder, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "WAIT_PAY").
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
		// 待核销订单数（今日创建）
		if res.TodayOrder.WaitVerifyOrder, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "WAIT_VERIFY").
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
		// 已完成订单数（今日创建）
		if res.TodayOrder.DoneOrder, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "DONE").
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
		// 已退款订单数（今日创建）
		if res.TodayOrder.RefundOrder, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderStatus, "REFUND").
			Where("DATE(created_at) = ?", today).
			Count(); err != nil {
			return
		}
	}
	return
}
