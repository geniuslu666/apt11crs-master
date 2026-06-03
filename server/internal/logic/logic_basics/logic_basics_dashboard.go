package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

type sBasicsDashboard struct{}

func NewBasicsDashboard() *sBasicsDashboard {
	return &sBasicsDashboard{}
}

func init() {
	service.RegisterBasicsDashboard(NewBasicsDashboard())
}

func (s *sBasicsDashboard) GetBasics(ctx context.Context, in *input_basics.BasicDashboardInp) (res *input_basics.BasicDashboardModel, err error) {

	res = &input_basics.BasicDashboardModel{}

	// 今日
	todayStart := gtime.Now().StartOfDay()
	todayEnd := gtime.Now().EndOfDay()
	// 1天前
	lastOneDayStart := gtime.New(todayStart).Add(time.Duration(-24) * time.Hour)
	lastOneDayEnd := gtime.New(todayEnd).Add(time.Duration(-24) * time.Hour)

	// 今日新增会员数=============
	res.MemberStat.TodayRegMemberNum, err = dao.PmsMember.Ctx(ctx).WhereBetween(dao.PmsMember.Columns().RegisterTime, todayStart, todayEnd).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	lastOneDayRegMemberNum, err := dao.PmsMember.Ctx(ctx).WhereBetween(dao.PmsMember.Columns().RegisterTime, lastOneDayStart, lastOneDayEnd).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日会员新增同比
	if lastOneDayRegMemberNum > 0 {
		res.MemberStat.TodayRegGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(res.MemberStat.TodayRegMemberNum-lastOneDayRegMemberNum)/float64(lastOneDayRegMemberNum)*100), 64)
	} else {
		res.MemberStat.TodayRegGrewPer = 0
		if res.MemberStat.TodayRegMemberNum > 0 {
			res.MemberStat.TodayRegGrewPer = 100
		}
	}

	// 累计会员============
	res.MemberStat.TotalMemberNum, err = dao.PmsMember.Ctx(ctx).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	beforeLastOneDayRegMemberNum, err := dao.PmsMember.Ctx(ctx).WhereLTE(dao.PmsMember.Columns().RegisterTime, lastOneDayEnd).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 累计会员新增同比
	if beforeLastOneDayRegMemberNum > 0 {
		res.MemberStat.TotalMemberGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(res.MemberStat.TotalMemberNum-beforeLastOneDayRegMemberNum)/float64(beforeLastOneDayRegMemberNum)*100), 64)
	} else {
		res.MemberStat.TotalMemberGrewPer = 0
		if res.MemberStat.TodayRegMemberNum > 0 {
			res.MemberStat.TotalMemberGrewPer = 100
		}
	}

	// 今日新增积分======================
	res.MemberStat.TodayIncBal, err = dao.PmsBalanceChange.Ctx(ctx).
		WhereGTE(dao.PmsBalanceChange.Columns().ChangePrice, 0).
		WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, todayStart, todayEnd).
		Sum(dao.PmsBalanceChange.Columns().ChangePrice)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	lastOneDayIncBal, err := dao.PmsBalanceChange.Ctx(ctx).
		WhereGTE(dao.PmsBalanceChange.Columns().ChangePrice, 0).
		WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Sum(dao.PmsBalanceChange.Columns().ChangePrice)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日积分新增同比
	if lastOneDayIncBal > 0 {
		res.MemberStat.TodayIncBalGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.MemberStat.TodayIncBal-lastOneDayIncBal)/lastOneDayIncBal*100), 64)
	} else {
		res.MemberStat.TodayIncBalGrewPer = 0
		if res.MemberStat.TodayIncBal > 0 {
			res.MemberStat.TodayIncBalGrewPer = 100
		}
	}
	// 今日消耗积分===================
	todayConsumeBal, err := dao.PmsBalanceChange.Ctx(ctx).
		WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0).
		WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, todayStart, todayEnd).
		Sum(dao.PmsBalanceChange.Columns().ChangePrice)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.MemberStat.TodayConsumeBal = math.Abs(todayConsumeBal)
	lastOneDayConsumeBal, err := dao.PmsBalanceChange.Ctx(ctx).
		WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0).
		WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Sum(dao.PmsBalanceChange.Columns().ChangePrice)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日消耗积分新增同比
	if math.Abs(lastOneDayConsumeBal) > 0 {
		res.MemberStat.TodayConsumeBalGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.MemberStat.TodayConsumeBal-lastOneDayConsumeBal)/math.Abs(lastOneDayConsumeBal)*100), 64)
	} else {
		if res.MemberStat.TodayConsumeBal > 0 {
			res.MemberStat.TodayConsumeBalGrewPer = 100
		} else {
			res.MemberStat.TodayConsumeBalGrewPer = 0
		}
	}

	// 今日订单数===================
	res.OrderStat.AllStat.TodayOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日订单数
	lastOneDayOrderNum, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日订单数新增同比
	if lastOneDayOrderNum > 0 {
		res.OrderStat.AllStat.TodayOrderNumGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TodayOrderNum-lastOneDayOrderNum)/lastOneDayOrderNum*100), 64)
	} else {
		if res.OrderStat.AllStat.TodayOrderNum > 0 {
			res.OrderStat.AllStat.TodayOrderNumGrewPer = 100
		} else {
			res.OrderStat.AllStat.TodayOrderNumGrewPer = 0
		}
	}
	// 全量订单 ================
	res.OrderStat.AllStat.TotalOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	lastOneDayBeforeOrderNum, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereLTE(dao.PmsAppStay.Columns().CreatedAt, lastOneDayEnd).
		Count()
	// 全量订单数新增同比
	if lastOneDayBeforeOrderNum > 0 {
		res.OrderStat.AllStat.TotalOrderNumGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(res.OrderStat.AllStat.TotalOrderNum-lastOneDayBeforeOrderNum)/float64(lastOneDayBeforeOrderNum)*100), 64)
	} else {
		if res.OrderStat.AllStat.TotalOrderNum > 0 {
			res.OrderStat.AllStat.TotalOrderNumGrewPer = 100
		} else {
			res.OrderStat.AllStat.TotalOrderNumGrewPer = 0
		}
	}
	// 今日退款单数===================
	res.OrderStat.AllStat.TodayRefundNum, err = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, todayStart, todayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日退款单数
	lastOneDayRefundNum, err := dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, lastOneDayStart, lastOneDayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日退款单数新增同比 (今日退款单数 - 昨日退款单数) / 昨日退款单数 * 100
	if lastOneDayRefundNum > 0 {
		res.OrderStat.AllStat.TodayRefundGrewPer = decimal.NewFromInt(gvar.New(res.OrderStat.AllStat.TodayRefundNum - lastOneDayRefundNum).Int64()).
			Div(decimal.NewFromInt(gvar.New(lastOneDayRefundNum).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).InexactFloat64()
	} else {
		if res.OrderStat.AllStat.TodayRefundNum > 0 {
			res.OrderStat.AllStat.TodayRefundGrewPer = 100
		} else {
			res.OrderStat.AllStat.TodayRefundGrewPer = 0
		}
	}

	// 全量退款单数===================
	res.OrderStat.AllStat.TotalRefundNum, err = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日前退款总单数
	lastOneDayBeforeRefundNum, err := dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereLTE(dao.PmsTransactionRefund.Columns().RefundTime, lastOneDayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 全量退款单数新增同比
	if lastOneDayBeforeRefundNum > 0 {
		res.OrderStat.AllStat.TotalRefundGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TotalRefundNum-lastOneDayBeforeRefundNum)/lastOneDayBeforeRefundNum*100), 64)
	} else {
		if res.OrderStat.AllStat.TotalRefundNum > 0 {
			res.OrderStat.AllStat.TotalRefundGrewPer = 100
		} else {
			res.OrderStat.AllStat.TotalRefundGrewPer = 0
		}
	}
	// 今日订单总额===================
	res.OrderStat.AllStat.TodayOrderMoney, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日订单金额
	lastOneDayOrderMoney, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日订单金额新增同比
	if lastOneDayOrderMoney > 0 {
		res.OrderStat.AllStat.TodayOrderMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TodayOrderMoney-lastOneDayOrderMoney)/lastOneDayOrderMoney*100), 64)
	} else {
		if res.OrderStat.AllStat.TodayOrderMoney > 0 {
			res.OrderStat.AllStat.TodayOrderMoneyGrewPer = 100
		} else {
			res.OrderStat.AllStat.TodayOrderMoneyGrewPer = 0
		}
	}

	// 全量订单总额===================
	res.OrderStat.AllStat.TotalOrderMoney, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日前订单总金额
	lastOneDayBeforeOrderMoney, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereLTE(dao.PmsAppStay.Columns().CreatedAt, lastOneDayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 总订单金额新增同比
	if lastOneDayBeforeOrderMoney > 0 {
		res.OrderStat.AllStat.TotalOrderMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TotalOrderMoney-lastOneDayBeforeOrderMoney)/lastOneDayBeforeOrderMoney*100), 64)
	} else {
		if res.OrderStat.AllStat.TotalOrderMoney > 0 {
			res.OrderStat.AllStat.TotalOrderMoneyGrewPer = 100
		} else {
			res.OrderStat.AllStat.TotalOrderMoneyGrewPer = 0
		}
	}

	// 今日退款总额===================
	res.OrderStat.AllStat.TodayRefundMoney, err = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, todayStart, todayEnd).
		Sum(dao.PmsTransactionRefund.Columns().RefundAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日订单退款总金额
	lastOneDayRefundMoney, err := dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, lastOneDayStart, lastOneDayEnd).
		Sum(dao.PmsTransactionRefund.Columns().RefundAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日退款总额新增同比
	if lastOneDayRefundMoney > 0 {
		res.OrderStat.AllStat.TodayRefundMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TodayRefundMoney-lastOneDayRefundMoney)/lastOneDayRefundMoney*100), 64)
	} else {
		if res.OrderStat.AllStat.TodayRefundMoney > 0 {
			res.OrderStat.AllStat.TodayRefundMoneyGrewPer = 100
		} else {
			res.OrderStat.AllStat.TodayRefundMoneyGrewPer = 0
		}
	}
	// 全量退款总额===================
	res.OrderStat.AllStat.TotalRefundMoney, err = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		Sum(dao.PmsTransactionRefund.Columns().RefundAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日前订单退款总金额
	lastOneDayBeforeRefundMoney, err := dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		WhereLTE(dao.PmsTransactionRefund.Columns().RefundTime, lastOneDayEnd).
		Sum(dao.PmsTransactionRefund.Columns().RefundAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 总退款金额新增同比
	if lastOneDayBeforeRefundMoney > 0 {
		res.OrderStat.AllStat.TotalRefundMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.AllStat.TotalRefundMoney-lastOneDayBeforeRefundMoney)/lastOneDayBeforeRefundMoney*100), 64)
	} else {
		if res.OrderStat.AllStat.TotalRefundMoney > 0 {
			res.OrderStat.AllStat.TotalRefundMoneyGrewPer = 100
		} else {
			res.OrderStat.AllStat.TotalRefundMoneyGrewPer = 0
		}
	}
	// 全部物业===================
	res.OrderStat.AllStat.TotalPropertyNum, err = dao.PmsProperty.Ctx(ctx).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 在线物业
	res.OrderStat.AllStat.OnlinePropertyNum, err = dao.PmsProperty.Ctx(ctx).Where(g.Map{
		dao.PmsProperty.Columns().Close:        1,
		dao.PmsProperty.Columns().BookingClose: 1,
	}).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}

	// 全部房型===================
	res.OrderStat.AllStat.TotalRoomTypeNum, err = dao.PmsRoomType.Ctx(ctx).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 在线房型
	propertyColumns, err := dao.PmsProperty.Ctx(ctx).Fields("uid").Where(g.Map{
		dao.PmsProperty.Columns().Close:        1,
		dao.PmsProperty.Columns().BookingClose: 1,
	}).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取uid失败！")
		return
	}
	uids := g.NewVar(propertyColumns).Strings()
	res.OrderStat.AllStat.OnlineRoomTypeNum, err = dao.PmsRoomType.Ctx(ctx).WhereIn(dao.PmsRoomType.Columns().Puid, uids).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}

	//----------app-------
	// 今日订单数===================
	res.OrderStat.HotelStat.TodayOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日订单数
	appLastOneDayOrderNum, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日订单数新增同比
	if appLastOneDayOrderNum > 0 {
		res.OrderStat.HotelStat.TodayOrderNumGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.HotelStat.TodayOrderNum-appLastOneDayOrderNum)/appLastOneDayOrderNum*100), 64)
	} else {
		if res.OrderStat.HotelStat.TodayOrderNum > 0 {
			res.OrderStat.HotelStat.TodayOrderNumGrewPer = 100
		} else {
			res.OrderStat.HotelStat.TodayOrderNumGrewPer = 0
		}
	}
	// 全量订单 ================
	res.OrderStat.HotelStat.TotalOrderNum, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	appLastOneDayBeforeOrderNum, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereLTE(dao.PmsAppStay.Columns().CreatedAt, lastOneDayEnd).
		Count()
	// 全量订单数新增同比
	if appLastOneDayBeforeOrderNum > 0 {
		res.OrderStat.HotelStat.TotalOrderNumGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(res.OrderStat.HotelStat.TotalOrderNum-appLastOneDayBeforeOrderNum)/float64(appLastOneDayBeforeOrderNum)*100), 64)
	} else {
		if res.OrderStat.HotelStat.TotalOrderNum > 0 {
			res.OrderStat.HotelStat.TotalOrderNumGrewPer = 100
		} else {
			res.OrderStat.HotelStat.TotalOrderNumGrewPer = 0
		}
	}

	// 今日退款单数===================
	res.OrderStat.HotelStat.TodayRefundNum = res.OrderStat.AllStat.TodayRefundNum

	// 今日退款单数新增同比 (今日退款单数 - 昨日退款单数) / 昨日退款单数 * 100
	res.OrderStat.HotelStat.TodayRefundGrewPer = res.OrderStat.AllStat.TodayRefundGrewPer

	// 全量退款单数===================
	res.OrderStat.HotelStat.TotalRefundNum = res.OrderStat.AllStat.TotalRefundNum
	res.OrderStat.HotelStat.TotalRefundGrewPer = res.OrderStat.AllStat.TotalRefundGrewPer

	// 今日订单总额===================
	res.OrderStat.HotelStat.TodayOrderMoney, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日订单金额
	appLastOneDayOrderMoney, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 今日订单金额新增同比
	if appLastOneDayOrderMoney > 0 {
		res.OrderStat.HotelStat.TodayOrderMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.HotelStat.TodayOrderMoney-appLastOneDayOrderMoney)/appLastOneDayOrderMoney*100), 64)
	} else {
		if res.OrderStat.HotelStat.TodayOrderMoney > 0 {
			res.OrderStat.HotelStat.TodayOrderMoneyGrewPer = 100
		} else {
			res.OrderStat.HotelStat.TodayOrderMoneyGrewPer = 0
		}
	}

	// 全量订单总额===================
	res.OrderStat.HotelStat.TotalOrderMoney, err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 昨日前订单总金额
	appLastOneDayBeforeOrderMoney, err := dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		WhereLTE(dao.PmsAppStay.Columns().CreatedAt, lastOneDayEnd).
		Sum(dao.PmsAppStay.Columns().OrderAmount)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	// 总订单金额新增同比
	if lastOneDayBeforeOrderMoney > 0 {
		res.OrderStat.HotelStat.TotalOrderMoneyGrewPer, err = strconv.ParseFloat(fmt.Sprintf("%.2f", (res.OrderStat.HotelStat.TotalOrderMoney-appLastOneDayBeforeOrderMoney)/appLastOneDayBeforeOrderMoney*100), 64)
	} else {
		if res.OrderStat.HotelStat.TotalOrderMoney > 0 {
			res.OrderStat.HotelStat.TotalOrderMoneyGrewPer = 100
		} else {
			res.OrderStat.HotelStat.TotalOrderMoneyGrewPer = 0
		}
	}

	// 今日退款总额===================
	res.OrderStat.HotelStat.TodayRefundMoney = res.OrderStat.AllStat.TodayRefundMoney

	// 今日退款总额新增同比
	res.OrderStat.HotelStat.TodayRefundMoneyGrewPer = res.OrderStat.AllStat.TodayRefundMoneyGrewPer

	// 全量退款总额===================
	res.OrderStat.HotelStat.TotalRefundMoney = res.OrderStat.AllStat.TotalRefundMoney

	// 总退款金额新增同比
	res.OrderStat.HotelStat.TotalRefundMoneyGrewPer = res.OrderStat.AllStat.TotalRefundMoneyGrewPer

	// 全部物业===================
	res.OrderStat.HotelStat.TotalPropertyNum = res.OrderStat.AllStat.TotalPropertyNum

	// 在线物业
	res.OrderStat.HotelStat.OnlinePropertyNum = res.OrderStat.AllStat.OnlinePropertyNum

	// 全部房型===================
	res.OrderStat.HotelStat.TotalRoomTypeNum = res.OrderStat.AllStat.TotalRoomTypeNum

	// 在线房型
	res.OrderStat.HotelStat.OnlineRoomTypeNum = res.OrderStat.AllStat.OnlineRoomTypeNum

	return
}

func (s *sBasicsDashboard) GetRanking(ctx context.Context, in *input_basics.RankingInp) (out *input_basics.RankingModel, err error) {
	var (
		PropertyALLNum    int64
		OTAChannelALLNum  int64
		NationalityALLNum int64
		DateList          []string
	)
	out = new(input_basics.RankingModel)

	// 查询三十天预订量
	if err = dao.PmsAppReservation.Ctx(ctx).
		Fields("DATE_FORMAT(created_at, '%Y-%m-%d') as orderDate, count(1) as orderNumber").
		Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
		Where(gdb.Raw("DATE_FORMAT(created_at, '%Y-%m-%d') > " + gtime.Now().Add(-30*24*time.Hour).Format("Y-m-d"))).
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		OrderDesc("DATE_FORMAT(created_at, '%Y-%m-%d')").Limit(30).Scan(&out.NewOrderNumberList); err != nil {
		return
	}

	// 计算七天的数组
	for i := 0; i < 7; i++ {
		DateList = append(DateList, gtime.Now().AddDate(0, 0, -i).Format("Y-m-d"))
	}
	// 查询七天预订量
	if err = dao.PmsAppReservation.Ctx(ctx).
		Fields("DATE_FORMAT(created_at, '%Y-%m-%d') as orderDate, count(1) as orderNumber").
		Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
		Where(gdb.Raw("DATE_FORMAT(created_at, '%Y-%m-%d') > " + gtime.Now().Add(-7*24*time.Hour).Format("Y-m-d"))).
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		OrderDesc("DATE_FORMAT(created_at, '%Y-%m-%d')").Limit(7).Scan(&out.OrderNumberList); err != nil {
		return
	}
outerLoop:
	for _, v := range DateList {
		for _, v2 := range out.OrderNumberList {
			if v == v2.OrderDate {
				continue outerLoop
			}
		}
		out.OrderNumberList = append(out.OrderNumberList, &input_basics.RankingOrderItem{OrderDate: v, OrderNumber: 0})
	}
	// 对 OrderNumberList 按照日期进行升序排序
	sort.Slice(out.OrderNumberList, func(i, j int) bool {
		return out.OrderNumberList[i].OrderDate < out.OrderNumberList[j].OrderDate
	})
	// 查询物业预订量
	if err = dao.PmsAppReservation.Ctx(ctx).
		WithAll().Hook(hook.PmsFindLanguageValueHook).
		Fields("COUNT(1) as number,puid").
		Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
		Group(dao.PmsAppStay.Columns().Puid).
		OrderDesc("COUNT(1)").Limit(5).
		Scan(&out.PropertyList); err != nil {
		return
	}
	for _, v := range out.PropertyList {
		PropertyALLNum += v.OrderNumber
	}
	for k, v := range out.PropertyList {
		out.PropertyList[k].Rate = decimal.NewFromInt(v.OrderNumber).
			Div(decimal.NewFromInt(PropertyALLNum)).
			Mul(decimal.NewFromInt(100)).
			Round(0).InexactFloat64()
	}

	if err = dao.PmsGuestProfile.Ctx(ctx).
		Fields("COUNT(1) as number,CASE WHEN nationality IN ('TW', 'HK', 'MO') THEN 'CN' ELSE nationality END as name").
		Group("CASE WHEN nationality IN ('TW', 'HK', 'MO') THEN 'CN' ELSE nationality END").
		OrderDesc("COUNT(1)").Limit(5).
		Scan(&out.NationalityList); err != nil {
		return
	}
	for _, v := range out.NationalityList {
		NationalityALLNum += v.OrderNumber
	}
	for k, v := range out.NationalityList {
		if g.IsEmpty(v.Name) {
			out.NationalityList[k].Name = "未知"
		}
		out.NationalityList[k].Rate = decimal.NewFromInt(v.OrderNumber).
			Div(decimal.NewFromInt(NationalityALLNum)).
			Mul(decimal.NewFromInt(100)).
			Round(0).InexactFloat64()
	}
	if err = dao.PmsAppReservation.Ctx(ctx).
		Fields("source_name as name,count(1) as number").
		Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
		WhereNotNull(dao.PmsAppReservation.Columns().SourceName).
		Group(dao.PmsAppReservation.Columns().SourceName).
		OrderDesc("COUNT(1)").Limit(5).
		Scan(&out.OTAChannelList); err != nil {
		return
	}
	for _, v := range out.OTAChannelList {
		OTAChannelALLNum += v.OrderNumber
	}
	for k, v := range out.OTAChannelList {
		out.OTAChannelList[k].Rate = decimal.NewFromInt(v.OrderNumber).
			Div(decimal.NewFromInt(OTAChannelALLNum)).
			Mul(decimal.NewFromInt(100)).
			Round(0).InexactFloat64()
	}
	return
}
