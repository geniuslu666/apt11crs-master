package admin

import (
	"APT/internal/dao"
	hook2 "APT/internal/library/hgorm/hook"
	"context"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) Dashboard(ctx context.Context, req *spa.DashboardReq) (res *spa.DashboardRes, err error) {

	res = new(spa.DashboardRes)
	if req.Type == "base" {
		//今日预约总数
		if res.Details.TodayOrderTotalNum, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//今日预约到店总数
		if res.Details.TodayDdOrderTotalNum, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().ServiceType, 1).
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//今日预约上门总数
		if res.Details.TodaySmOrderTotalNum, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().ServiceType, 2).
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}

		// 今日营业额
		if res.Details.TodayOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 总营业额
		if res.Details.TotalOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 全部预约数
		if res.Details.TotalOrderTotalNum, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}
		// 有效预约数
		if res.Details.EffectiveTotalOrderTotalNum, err = dao.SpaOrder.Ctx(ctx).
			WhereNot(dao.SpaOrder.Columns().OrderStatus, "CANCEL").
			Where(dao.SpaOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}
		// 空闲中技师数
		if res.Details.CanOrderTechnicianTotal, err = dao.SpaTechnician.Ctx(ctx).
			Where(dao.SpaTechnician.Columns().WorkStatus, "CANORDER").
			Count(); err != nil {
			return
		}

		// 工作中技师数
		if res.Details.WorkingTechnicianTotal, err = dao.SpaTechnician.Ctx(ctx).
			Where(dao.SpaTechnician.Columns().WorkStatus, "WORKING").
			Count(); err != nil {
			return
		}

		// 休息中技师数
		if res.Details.RestTechnicianTotal, err = dao.SpaTechnician.Ctx(ctx).
			Where(dao.SpaTechnician.Columns().WorkStatus, "REST").
			Count(); err != nil {
			return
		}

	} else if req.Type == "todayOrder" {
		//待付款订单
		if res.TodayOrder.WaitPayOrder, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_PAY").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待付款订单金额
		if res.TodayOrder.WaitPayOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_PAY").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		//待确认订单
		if res.TodayOrder.WaitConfirmOrder, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_CONFIRM").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待确认订单金额
		if res.TodayOrder.WaitConfirmOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_CONFIRM").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		//待服务订单
		if res.TodayOrder.WaitServeOrder, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_SERVE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待服务订单金额
		if res.TodayOrder.WaitServeOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_SERVE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		//服务中订单
		if res.TodayOrder.ServiceDoingOrder, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "SERVING").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//服务中订单金额
		if res.TodayOrder.ServiceDoingOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "SERVING").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}
		//服务完成订单
		if res.TodayOrder.ServiceCompleteOrder, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "DONE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//服务完成订单金额
		if res.TodayOrder.ServiceCompleteOrderAmount, err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderStatus, "DONE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.SpaOrder.Columns().OrderAmount); err != nil {
			return
		}

	} else if req.Type == "serviceLine" {
		if req.ServiceLineType == "num" {
			if err = dao.SpaService.Ctx(ctx).Where(dao.SpaService.Columns().ServiceState, 1).Page(1, 5).OrderDesc(dao.SpaService.Columns().TotalOrderNum).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.ServiceLine); err != nil {
				return
			}
		} else {
			if err = dao.SpaService.Ctx(ctx).Where(dao.SpaService.Columns().ServiceState, 1).Page(1, 5).OrderDesc(dao.SpaService.Columns().TotalOrderAmount).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.ServiceLine); err != nil {
				return
			}
		}
	} else if req.Type == "technicianLine" {
		if req.TechnicianLineType == "num" {
			if err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Status, 1).Page(1, 5).WithAll().OrderDesc(dao.SpaTechnician.Columns().PayOrderNum).Scan(&res.TechnicianLine); err != nil {
				return
			}
		} else {
			if err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Status, 1).Page(1, 5).WithAll().OrderDesc(dao.SpaTechnician.Columns().PayOrderAmount).Scan(&res.TechnicianLine); err != nil {
				return
			}
		}
	}
	return
}
