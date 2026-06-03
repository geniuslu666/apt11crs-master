package logic_car

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"APT/utility/uuid"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
)

type sCarSettlementOrder struct{}

func NewCarSettlementOrder() *sCarSettlementOrder {
	return &sCarSettlementOrder{}
}

func init() {
	service.RegisterCarSettlementOrder(NewCarSettlementOrder())
}

// Model 结算模式ORM模型
func (s *sCarSettlementOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarSettlementOrder.Ctx(ctx), option...)
}

// DailySettlement 每日结算
func (s *sCarSettlementOrder) DailySettlement(ctx context.Context) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartTime *gtime.Time
			EndTime   *gtime.Time
		)

		StartTime = gtime.New(gtime.Now().StartOfDay()).Add(time.Duration(-24) * time.Hour)
		EndTime = gtime.Now().StartOfDay()

		// 获取结算信息(周期性按日结算)
		var SettlementList []*input_car.CarOrderSettleInfoModel
		err = dao.CarOrder.Ctx(ctx).
			Fields(dao.CarOrder.Columns().DriverId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.CarOrder.Columns().OrderType, "CRS").
			Where(dao.CarOrder.Columns().SettlementType, 2).
			Where(dao.CarOrder.Columns().SettlementCycle, 1).
			Where(dao.CarOrder.Columns().OrderStatus, "DONE").
			WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
			WhereGT(dao.CarOrder.Columns().DriverId, 0).
			OmitEmptyWhere().
			Group(dao.CarOrder.Columns().DriverId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.CarSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("CS"),
				dao.CarSettlementOrder.Columns().DriverId:         gvar.New(SettlementInfo.DriverId).Int(),
				dao.CarSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.CarSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.CarSettlementOrder.Columns().Status:           "DONE",
				dao.CarSettlementOrder.Columns().StartTime:        StartTime,
				dao.CarSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_car.CarSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.CarOrder.Ctx(ctx).
				Fields(dao.CarOrder.Columns().Id).
				Where(dao.CarOrder.Columns().DriverId, SettlementInfo.DriverId).
				Where(dao.CarOrder.Columns().OrderType, "CRS").
				Where(dao.CarOrder.Columns().SettlementType, 2).
				Where(dao.CarOrder.Columns().SettlementCycle, 1).
				Where(dao.CarOrder.Columns().OrderStatus, "DONE").
				WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
				Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.CarOrder.Ctx(ctx).
				WhereIn(dao.CarOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.CarOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.CarOrder.Columns().SettlementTime:    gtime.Now(),
					dao.CarOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新司机预定量和预定金额
			if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, SettlementInfo.DriverId).Update(g.MapStrAny{
				dao.CarDriver.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.CarDriver.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.CarDriver.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				//dao.CarDriver.Columns().Balance:               gdb.Raw(fmt.Sprintf("balance+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
				return
			}

			// 写入司机积分变更日志
			/*if _, err = dao.CarDriverBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
				DriverId:          SettlementInfo.DriverId,
				Type:              "SETTLEMENT",
				ChangePrice:       SettlementInfo.TotalSettlementAmount,
				SettlementOrderId: int(SettlementOrderId),
				Des:               "订单收入",
				CreatedAt:         gtime.Now(),
				UpdatedAt:         gtime.Now(),
			}); err != nil {
				err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
				return
			}*/

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
					OrderStatus: "DONE",
					ActionWay:   "SETTLEMENT",
					Remark:      "订单已结算",
					OperateType: "SYSTEM",
				}); err != nil {
					return err
				}
			}
		}

		return
	})
}

// WeekSettlement 每周结算
func (s *sCarSettlementOrder) WeekSettlement(ctx context.Context) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartTime *gtime.Time
			EndTime   *gtime.Time
		)

		// 当前时间
		now := time.Now()
		// 获取当前时间的星期几
		weekday := now.Weekday()
		// 判断是否是周一
		if int(weekday) != 1 {
			return
		}

		StartTime = gtime.New(gtime.Now().StartOfDay()).Add(time.Duration(-24*7) * time.Hour)
		EndTime = gtime.Now().StartOfDay()

		// 获取结算信息
		var SettlementList []*input_car.CarOrderSettleInfoModel
		err = dao.CarOrder.Ctx(ctx).
			Fields(dao.CarOrder.Columns().DriverId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.CarOrder.Columns().OrderType, "CRS").
			Where(dao.CarOrder.Columns().SettlementType, 2).
			Where(dao.CarOrder.Columns().SettlementCycle, 2).
			Where(dao.CarOrder.Columns().OrderStatus, "DONE").
			WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
			WhereGT(dao.CarOrder.Columns().DriverId, 0).
			OmitEmptyWhere().
			Group(dao.CarOrder.Columns().DriverId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.CarSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("CS"),
				dao.CarSettlementOrder.Columns().DriverId:         gvar.New(SettlementInfo.DriverId).Int(),
				dao.CarSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.CarSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.CarSettlementOrder.Columns().Status:           "DONE",
				dao.CarSettlementOrder.Columns().StartTime:        StartTime,
				dao.CarSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_car.CarSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.CarOrder.Ctx(ctx).
				Fields(dao.CarOrder.Columns().Id).
				Where(dao.CarOrder.Columns().DriverId, SettlementInfo.DriverId).
				Where(dao.CarOrder.Columns().OrderType, "CRS").
				Where(dao.CarOrder.Columns().SettlementType, 2).
				Where(dao.CarOrder.Columns().SettlementCycle, 2).
				Where(dao.CarOrder.Columns().OrderStatus, "DONE").
				WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
				Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.CarOrder.Ctx(ctx).
				WhereIn(dao.CarOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.CarOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.CarOrder.Columns().SettlementTime:    gtime.Now(),
					dao.CarOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新司机预定量和预定金额
			if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, SettlementInfo.DriverId).Update(g.MapStrAny{
				dao.CarDriver.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.CarDriver.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.CarDriver.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				//dao.CarDriver.Columns().Balance:               gdb.Raw(fmt.Sprintf("balance+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
				return
			}

			// 写入司机积分变更日志
			/*if _, err = dao.CarDriverBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
				DriverId:          SettlementInfo.DriverId,
				Type:              "SETTLEMENT",
				ChangePrice:       SettlementInfo.TotalSettlementAmount,
				SettlementOrderId: int(SettlementOrderId),
				Des:               "订单收入",
				CreatedAt:         gtime.Now(),
				UpdatedAt:         gtime.Now(),
			}); err != nil {
				err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
				return
			}*/

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
					OrderStatus: "DONE",
					ActionWay:   "SETTLEMENT",
					Remark:      "订单已结算",
					OperateType: "SYSTEM",
				}); err != nil {
					return err
				}
			}
		}

		return
	})
}

// MonthSettlement 每月结算
func (s *sCarSettlementOrder) MonthSettlement(ctx context.Context) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartTime *gtime.Time
			EndTime   *gtime.Time
		)
		// 当前时间
		now := time.Now()
		// 获取当前月的第几天
		monthDay := now.Day()
		// 判断是否是1号
		if int(monthDay) != 1 {
			return
		}
		StartTime = gtime.New(gtime.New(gtime.Now().AddDate(0, -1, 0)).Format("Y-m") + "-01").StartOfDay()
		EndTime = StartTime.AddDate(0, 1, -1).EndOfDay()

		// 获取结算信息
		var SettlementList []*input_car.CarOrderSettleInfoModel
		err = dao.CarOrder.Ctx(ctx).
			Fields(dao.CarOrder.Columns().DriverId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.CarOrder.Columns().OrderType, "CRS").
			Where(dao.CarOrder.Columns().SettlementType, 2).
			Where(dao.CarOrder.Columns().SettlementCycle, 3).
			Where(dao.CarOrder.Columns().OrderStatus, "DONE").
			WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
			WhereGT(dao.CarOrder.Columns().DriverId, 0).
			OmitEmptyWhere().
			Group(dao.CarOrder.Columns().DriverId).
			Scan(&SettlementList)
		if err != nil {
			return
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.CarSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("CS"),
				dao.CarSettlementOrder.Columns().DriverId:         gvar.New(SettlementInfo.DriverId).Int(),
				dao.CarSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.CarSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.CarSettlementOrder.Columns().Status:           "DONE",
				dao.CarSettlementOrder.Columns().StartTime:        StartTime,
				dao.CarSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_car.CarSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.CarOrder.Ctx(ctx).
				Fields(dao.CarOrder.Columns().Id).
				Where(dao.CarOrder.Columns().DriverId, SettlementInfo.DriverId).
				Where(dao.CarOrder.Columns().OrderType, "CRS").
				Where(dao.CarOrder.Columns().SettlementType, 2).
				Where(dao.CarOrder.Columns().SettlementCycle, 3).
				Where(dao.CarOrder.Columns().OrderStatus, "DONE").
				WhereBetween(dao.CarOrder.Columns().ActualEndTime, StartTime, EndTime).
				Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.CarOrder.Ctx(ctx).
				WhereIn(dao.CarOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.CarOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.CarOrder.Columns().SettlementTime:    gtime.Now(),
					dao.CarOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新司机预定量和预定金额
			if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, SettlementInfo.DriverId).Update(g.MapStrAny{
				dao.CarDriver.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.CarDriver.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.CarDriver.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				//dao.CarDriver.Columns().Balance:               gdb.Raw(fmt.Sprintf("balance+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
				return
			}

			// 写入司机积分变更日志
			/*if _, err = dao.CarDriverBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
				DriverId:          SettlementInfo.DriverId,
				Type:              "SETTLEMENT",
				ChangePrice:       SettlementInfo.TotalSettlementAmount,
				SettlementOrderId: int(SettlementOrderId),
				Des:               "订单收入",
				CreatedAt:         gtime.Now(),
				UpdatedAt:         gtime.Now(),
			}); err != nil {
				err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
				return
			}*/

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
					OrderStatus: "DONE",
					ActionWay:   "SETTLEMENT",
					Remark:      "订单已结算",
					OperateType: "SYSTEM",
				}); err != nil {
					return err
				}
			}
		}

		return
	})

}

// SettlementOrderStat 结算概况
func (s *sCarSettlementOrder) SettlementOrderStat(ctx context.Context, in *input_car.CarSettlementOrderStatInp) (res *input_car.CarSettlementOrderStatModel, err error) {
	res = &input_car.CarSettlementOrderStatModel{}
	mod := dao.CarOrder.Ctx(ctx)

	// 订单总额(已支付)
	res.TotalOrderAmount, err = mod.WhereNotNull(dao.CarOrder.Columns().PayTime).Sum(dao.CarOrder.Columns().OrderAmount)

	// 待结算金额(已支付已完成但还未结算的金额)
	res.WaitSettlementAmount, err = mod.
		WhereNotNull(dao.CarOrder.Columns().PayTime).
		Where(dao.CarOrder.Columns().OrderStatus, "DONE").
		Where(dao.CarOrder.Columns().SettlementStatus, "WAIT").
		Sum(dao.CarOrder.Columns().SettlementAmount)

	// 已结算
	res.DoneSettlementAmount, err = mod.
		Where(dao.CarOrder.Columns().OrderStatus, "DONE").
		Where(dao.CarOrder.Columns().SettlementStatus, "SUCCESS").
		Sum(dao.CarOrder.Columns().SettlementAmount)

	// 已核账
	res.HadVerifySettlementAmount, err = dao.CarSettlementOrder.Ctx(ctx).
		Where(dao.CarSettlementOrder.Columns().VerifyStatus, "VERIFIED").
		Sum(dao.CarSettlementOrder.Columns().SettlementAmount)

	return
}

// List 结算列表
func (s *sCarSettlementOrder) List(ctx context.Context, in *input_car.CarSettlementOrderListInp) (list []*input_car.CarSettlementOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_car.CarSettlementOrderListModel{})

	if in.DriverId > 0 {
		mod = mod.Where(dao.CarSettlementOrder.Columns().DriverId, in.DriverId)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.CarSettlementOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.VerifyStatus) {
		mod = mod.Where(dao.CarSettlementOrder.Columns().VerifyStatus, in.VerifyStatus)
	}

	if len(in.VerifyTime) == 2 {
		mod = mod.WhereBetween(dao.CarSettlementOrder.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.CarSettlementOrder.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算单列表失败，请稍后重试！")
		return
	}
	return
}

// View 结算单详情
func (s *sCarSettlementOrder) View(ctx context.Context, in *input_car.CarSettlementOrderViewInp) (res *input_car.CarSettlementOrderViewModel, err error) {
	if err = s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算单详情失败，请稍后重试！")
		return
	}
	return
}

// Verify 核账
func (s *sCarSettlementOrder) Verify(ctx context.Context, in *input_car.CarSettlementOrderVerifyInp) (err error) {

	var models *entity.CarSettlementOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("结算信息不存在或已被删除")
		return
	}

	if models.VerifyStatus != "WAIT_VERIFY" {
		err = gerror.New("结算单状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 改变结算单状态
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarSettlementOrderVerifyFields{
			VerifyStatus:    "VERIFIED",
			VerifyImg:       in.VerifyImg,
			VerifyDesc:      in.VerifyDesc,
			VerifyTime:      gtime.Now(),
			VerifyOperateId: contexts.GetUserId(ctx),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 更新司机预定量和预定金额
		if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, models.DriverId).Update(g.MapStrAny{
			dao.CarDriver.Columns().VerifyMoney: gdb.Raw(fmt.Sprintf("verify_money+%f", models.SettlementAmount)),
			dao.CarDriver.Columns().Balance:     gdb.Raw(fmt.Sprintf("balance+%f", models.SettlementAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
			return
		}

		// 写入司机积分变更日志
		if _, err = dao.CarDriverBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
			DriverId:          models.DriverId,
			Type:              "SETTLEMENT",
			ChangePrice:       models.SettlementAmount,
			SettlementOrderId: int(models.Id),
			Des:               "订单收入",
			CreatedAt:         gtime.Now(),
			UpdatedAt:         gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
			return
		}

		/*if _, err = dao.CarDriverBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
			DriverId:          models.DriverId,
			Type:              "VERIFY",
			ChangePrice:       -models.SettlementAmount,
			SettlementOrderId: int(models.Id),
			Des:               "核账",
			CreatedAt:         gtime.Now(),
			UpdatedAt:         gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
			return
		}*/

		OrderIds, _ := dao.CarOrder.Ctx(ctx).
			Fields(dao.CarOrder.Columns().Id).
			Where(dao.CarOrder.Columns().SettlementOrderId, in.Id).
			Array()

		// 订单日志
		for _, OrderId := range OrderIds {
			if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
				OrderId:     g.NewVar(OrderId).Int(),
				OrderStatus: "DONE",
				ActionWay:   "SETTLEMENT_VERIFY",
				Remark:      "订单已核账",
				OperateType: "ADMIN",
				OperateId:   int(contexts.GetUserId(ctx)),
			}); err != nil {
				return err
			}
		}
		return
	})

}
