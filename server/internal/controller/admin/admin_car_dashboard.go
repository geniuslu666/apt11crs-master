package admin

import (
	"APT/api/admin/car"
	"APT/internal/dao"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerCar) Dashboard(ctx context.Context, req *car.DashboardReq) (res *car.DashboardRes, err error) {
	res = new(car.DashboardRes)
	if req.Type == "base" {
		//今日预约总数
		if res.Details.TodayOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//今日预约接机总数
		if res.Details.TodayPickUpOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "PICKUP").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//今日预约接机送机总数
		if res.Details.TodayDeliveryOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "DELIVERY").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//今日预约包车送机总数
		if res.Details.TodayCharteredCarOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "CHARTERED").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}

		// 今日营业额
		if res.Details.TodayOrderAmount, err = dao.CarOrder.Ctx(ctx).
			WhereIn(dao.CarOrder.Columns().PayStatus, g.Slice{"HAVE_PAID", "REFUND"}).
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 今日退款额
		if res.Details.TodayOrderRefund, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().PayStatus, "REFUND").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}

		// 全部累计预约数
		if res.Details.TotalOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			WhereIn(dao.CarOrder.Columns().PayStatus, g.Slice{"HAVE_PAID", "REFUND"}).
			Count(); err != nil {
			return
		}
		// 有效累计预约数
		if res.Details.EffectiveOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}
		// 全部预约接机总数
		if res.Details.TotalPickUpOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "PICKUP").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}
		// 全部预约送机总数
		if res.Details.TotalDeliveryOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "DELIVERY").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}
		// 全部预约包车总数
		if res.Details.TotalCharteredCarOrderTotalNum, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().ServiceType, "CHARTERED").
			Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").
			Count(); err != nil {
			return
		}

		// 累计营业额
		if res.Details.TotalOrderAmount, err = dao.CarOrder.Ctx(ctx).
			WhereIn(dao.CarOrder.Columns().PayStatus, g.Slice{"HAVE_PAID", "REFUND"}).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		// 累计退款额
		if res.Details.TotalOrderRefund, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().PayStatus, "REFUND").
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}

		// 司机总数
		if res.Details.TotalDriverNum, err = dao.CarDriver.Ctx(ctx).
			Count(); err != nil {
			return
		}
		// 正常司机数
		if res.Details.OnDriverTotal, err = dao.CarDriver.Ctx(ctx).
			Where(dao.CarDriver.Columns().Status, 1).
			Count(); err != nil {
			return
		}
		// 休息中司机数
		if res.Details.RestDriverTotal, err = dao.CarDriver.Ctx(ctx).
			Where(dao.CarDriver.Columns().WorkStatus, "REST").
			Where(dao.CarDriver.Columns().Status, 1).
			Count(); err != nil {
			return
		}
		// 停用司机数
		if res.Details.OffDriverTotal, err = dao.CarDriver.Ctx(ctx).
			Where(dao.CarDriver.Columns().Status, 2).
			Count(); err != nil {
			return
		}

		// 车辆总数
		if res.Details.TotalCarNum, err = dao.CarCar.Ctx(ctx).
			Count(); err != nil {
			return
		}
		// 正常车辆数
		if res.Details.OnCarTotal, err = dao.CarCar.Ctx(ctx).
			Where(dao.CarCar.Columns().Status, 1).
			Count(); err != nil {
			return
		}
		// 停用车辆数
		if res.Details.OffCarTotal, err = dao.CarCar.Ctx(ctx).
			Where(dao.CarCar.Columns().Status, 2).
			Count(); err != nil {
			return
		}

	} else if req.Type == "todayOrder" {
		//待付款订单
		if res.TodayOrder.WaitPayOrder, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_PAY").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待付款订单金额
		if res.TodayOrder.WaitPayOrderAmount, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_PAY").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		//待确认订单
		if res.TodayOrder.WaitConfirmOrder, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_CONFIRM").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待确认订单金额
		if res.TodayOrder.WaitConfirmOrderAmount, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_CONFIRM").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		//待服务订单
		if res.TodayOrder.WaitServeOrder, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_SERVE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//待服务订单金额
		if res.TodayOrder.WaitServeOrderAmount, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "WAIT_SERVE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		//服务中订单
		if res.TodayOrder.ServiceDoingOrder, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "SERVING").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//服务中订单金额
		if res.TodayOrder.ServiceDoingOrderAmount, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "SERVING").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}
		//服务完成订单
		if res.TodayOrder.ServiceCompleteOrder, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "DONE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Count(); err != nil {
			return
		}
		//服务完成订单金额
		if res.TodayOrder.ServiceCompleteOrderAmount, err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderStatus, "DONE").
			Where("DATE(created_at) = ?", gtime.Now().Format("Y-m-d")).
			Sum(dao.CarOrder.Columns().OrderAmount); err != nil {
			return
		}

	} else if req.Type == "serviceLine" {
		if req.ServiceLineType == "num" {
			if err = dao.CarService.Ctx(ctx).Where(dao.CarService.Columns().Status, 1).Page(1, 5).Fields(dao.CarService.Columns().ServiceType).Fields(gdb.Raw("SUM( total_order_num ) as totalOrderNum")).Group(dao.CarService.Columns().ServiceType).OrderDesc("totalOrderNum").Scan(&res.ServiceLine); err != nil {
				return
			}
		} else {
			if err = dao.CarService.Ctx(ctx).Where(dao.CarService.Columns().Status, 1).Page(1, 5).Fields(dao.CarService.Columns().ServiceType).Fields(gdb.Raw("SUM( total_order_amount ) as totalOrderAmount")).Group(dao.CarService.Columns().ServiceType).OrderDesc("totalOrderAmount").Scan(&res.ServiceLine); err != nil {
				return
			}
		}
	} else if req.Type == "driverLine" {
		if req.DriverLineType == "num" {
			if err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Status, 1).Page(1, 5).WithAll().OrderDesc(dao.CarDriver.Columns().PayOrderNum).Scan(&res.DriverLine); err != nil {
				return
			}
		} else {
			if err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Status, 1).Page(1, 5).WithAll().OrderDesc(dao.CarDriver.Columns().PayOrderAmount).Scan(&res.DriverLine); err != nil {
				return
			}
		}
	}
	return
}
