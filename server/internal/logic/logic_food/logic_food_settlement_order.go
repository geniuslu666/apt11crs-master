package logic_food

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"APT/utility/uuid"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodSettlementOrder struct{}

func NewFoodSettlementOrder() *sFoodSettlementOrder {
	return &sFoodSettlementOrder{}
}

func init() {
	service.RegisterFoodSettlementOrder(NewFoodSettlementOrder())
}

func (s *sFoodSettlementOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodSettlementOrder.Ctx(ctx), option...)
}

// DailySettlement 每日结算
func (s *sFoodSettlementOrder) DailySettlement(ctx context.Context) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartTime *gtime.Time
			EndTime   *gtime.Time
		)

		StartTime = gtime.New(gtime.Now().StartOfDay()).Add(time.Duration(-24) * time.Hour)
		EndTime = gtime.Now().StartOfDay()

		/*sql, err := gdb.ToSQL(ctx, func(ctx context.Context) error {
			_, err = dao.FoodOrder.Ctx(ctx).
				Fields(dao.FoodOrder.Columns().RestaurantId).
				Fields(gdb.Raw("SUM( settlement_amount ) as total_settlement_amount")).
				Fields(gdb.Raw("SUM( order_amount ) as total_order_amount")).
				Where(dao.FoodOrder.Columns().SettlementType, 2).
				Where(dao.FoodOrder.Columns().SettlementCycle, 1).
				Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
				WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
				Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
				OmitEmptyWhere().
				Group(dao.FoodOrder.Columns().RestaurantId).
				All()
			return err
		})

		fmt.Println("-------------SQL---------------")
		fmt.Println(sql)*/

		// 获取结算信息
		var SettlementList []*input_food.FoodOrderSettleInfoModel
		err = dao.FoodOrder.Ctx(ctx).
			Fields(dao.FoodOrder.Columns().RestaurantId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.FoodOrder.Columns().SettlementType, 2).
			Where(dao.FoodOrder.Columns().SettlementCycle, 1).
			Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
			WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
			Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
			OmitEmptyWhere().
			Group(dao.FoodOrder.Columns().RestaurantId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.FoodSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("FS"),
				dao.FoodSettlementOrder.Columns().RestaurantId:     gvar.New(SettlementInfo.RestaurantId).Int(),
				dao.FoodSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.FoodSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.FoodSettlementOrder.Columns().Status:           "DONE",
				dao.FoodSettlementOrder.Columns().StartTime:        StartTime,
				dao.FoodSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_food.FoodSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.FoodOrder.Ctx(ctx).
				Fields(dao.FoodOrder.Columns().Id).
				Where(dao.FoodOrder.Columns().RestaurantId, SettlementInfo.RestaurantId).
				Where(dao.FoodOrder.Columns().SettlementType, 2).
				Where(dao.FoodOrder.Columns().SettlementCycle, 1).
				Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
				WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
				Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.FoodOrder.Ctx(ctx).
				WhereIn(dao.FoodOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.FoodOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.FoodOrder.Columns().SettlementTime:    gtime.Now(),
					dao.FoodOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新餐厅结算统计信息
			if _, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, SettlementInfo.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.FoodRestaurant.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.FoodRestaurant.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
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
func (s *sFoodSettlementOrder) WeekSettlement(ctx context.Context) (err error) {
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
		var SettlementList []*input_food.FoodOrderSettleInfoModel
		err = dao.FoodOrder.Ctx(ctx).
			Fields(dao.FoodOrder.Columns().RestaurantId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.FoodOrder.Columns().SettlementType, 2).
			Where(dao.FoodOrder.Columns().SettlementCycle, 2).
			Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
			WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
			Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
			OmitEmptyWhere().
			Group(dao.FoodOrder.Columns().RestaurantId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.FoodSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("FS"),
				dao.FoodSettlementOrder.Columns().RestaurantId:     gvar.New(SettlementInfo.RestaurantId).Int(),
				dao.FoodSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.FoodSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.FoodSettlementOrder.Columns().Status:           "DONE",
				dao.FoodSettlementOrder.Columns().StartTime:        StartTime,
				dao.FoodSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_food.FoodSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.FoodOrder.Ctx(ctx).
				Fields(dao.FoodOrder.Columns().Id).
				Where(dao.FoodOrder.Columns().RestaurantId, SettlementInfo.RestaurantId).
				Where(dao.FoodOrder.Columns().SettlementType, 2).
				Where(dao.FoodOrder.Columns().SettlementCycle, 2).
				Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
				WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
				Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.FoodOrder.Ctx(ctx).
				WhereIn(dao.FoodOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.FoodOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.FoodOrder.Columns().SettlementTime:    gtime.Now(),
					dao.FoodOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新餐厅结算统计信息
			if _, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, SettlementInfo.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.FoodRestaurant.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.FoodRestaurant.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
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
func (s *sFoodSettlementOrder) MonthSettlement(ctx context.Context) (err error) {
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
		var SettlementList []*input_food.FoodOrderSettleInfoModel
		err = dao.FoodOrder.Ctx(ctx).
			Fields(dao.FoodOrder.Columns().RestaurantId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( id ) as totalOrderNum")).
			Where(dao.FoodOrder.Columns().SettlementType, 2).
			Where(dao.FoodOrder.Columns().SettlementCycle, 3).
			Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
			WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
			Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
			OmitEmptyWhere().
			Group(dao.FoodOrder.Columns().RestaurantId).
			Scan(&SettlementList)
		if err != nil {
			return
		}

		for _, SettlementInfo := range SettlementList {
			InsertSettlementMap := g.MapStrAny{
				dao.FoodSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("FS"),
				dao.FoodSettlementOrder.Columns().RestaurantId:     gvar.New(SettlementInfo.RestaurantId).Int(),
				dao.FoodSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.FoodSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.FoodSettlementOrder.Columns().Status:           "DONE",
				dao.FoodSettlementOrder.Columns().StartTime:        StartTime,
				dao.FoodSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_food.FoodSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 修改订单结算状态
			OrderIds, _ := dao.FoodOrder.Ctx(ctx).
				Fields(dao.FoodOrder.Columns().Id).
				Where(dao.FoodOrder.Columns().RestaurantId, SettlementInfo.RestaurantId).
				Where(dao.FoodOrder.Columns().SettlementType, 2).
				Where(dao.FoodOrder.Columns().SettlementCycle, 3).
				Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
				WhereBetween(dao.FoodOrder.Columns().VerifyTime, StartTime, EndTime).
				Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
				Array()

			if _, err = dao.FoodOrder.Ctx(ctx).
				WhereIn(dao.FoodOrder.Columns().Id, OrderIds).
				Data(g.MapStrAny{
					dao.FoodOrder.Columns().SettlementStatus:  "SUCCESS",
					dao.FoodOrder.Columns().SettlementTime:    gtime.Now(),
					dao.FoodOrder.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单结算状态失败，请稍后重试！")
			}

			// 更新餐厅结算统计信息
			if _, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, SettlementInfo.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
				dao.FoodRestaurant.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
				dao.FoodRestaurant.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}

			// 订单日志
			for _, OrderId := range OrderIds {
				if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
					OrderId:     g.NewVar(OrderId).Int(),
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
func (s *sFoodSettlementOrder) SettlementOrderStat(ctx context.Context, in *input_food.FoodSettlementOrderStatInp) (res *input_food.FoodSettlementOrderStatModel, err error) {
	res = &input_food.FoodSettlementOrderStatModel{}
	mod := dao.FoodOrder.Ctx(ctx)

	// 订单总额(已支付)
	res.TotalOrderAmount, err = mod.WhereNotNull(dao.FoodOrder.Columns().PayTime).Where(dao.FoodOrder.Columns().RefundStatus, "WAIT").Sum(dao.FoodOrder.Columns().OrderAmount)

	// 待结算金额(已支付已核销但还未结算的金额)
	res.WaitSettlementAmount, err = mod.
		WhereNotNull(dao.FoodOrder.Columns().PayTime).
		Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
		Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
		Sum(dao.FoodOrder.Columns().SettlementAmount)

	// 已结算
	res.DoneSettlementAmount, err = mod.
		Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
		Where(dao.FoodOrder.Columns().SettlementStatus, "SUCCESS").
		Sum(dao.FoodOrder.Columns().SettlementAmount)

	// 已核账
	res.HadVerifySettlementAmount, err = dao.FoodSettlementOrder.Ctx(ctx).
		Where(dao.FoodSettlementOrder.Columns().VerifyStatus, "VERIFIED").
		Sum(dao.FoodSettlementOrder.Columns().SettlementAmount)

	return
}

// List 结算列表
func (s *sFoodSettlementOrder) List(ctx context.Context, in *input_food.FoodSettlementOrderListInp) (list []*input_food.FoodSettlementOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_food.FoodSettlementOrderListModel{})

	if in.RestaurantId > 0 {
		mod = mod.Where(dao.FoodSettlementOrder.Columns().RestaurantId, in.RestaurantId)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.FoodSettlementOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.VerifyStatus) {
		mod = mod.Where(dao.FoodSettlementOrder.Columns().VerifyStatus, in.VerifyStatus)
	}

	if len(in.VerifyTime) == 2 {
		mod = mod.WhereBetween(dao.FoodSettlementOrder.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.FoodSettlementOrder.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算单列表失败，请稍后重试！")
		return
	}
	return
}

// View 结算单详情
func (s *sFoodSettlementOrder) View(ctx context.Context, in *input_food.FoodSettlementOrderViewInp) (res *input_food.FoodSettlementOrderViewModel, err error) {
	if err = s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算单详情失败，请稍后重试！")
		return
	}
	return
}

// Verify 核账
func (s *sFoodSettlementOrder) Verify(ctx context.Context, in *input_food.FoodSettlementOrderVerifyInp) (err error) {

	var models *entity.FoodSettlementOrder
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
			WherePri(in.Id).Data(input_food.FoodSettlementOrderVerifyFields{
			VerifyStatus:    "VERIFIED",
			VerifyImg:       in.VerifyImg,
			VerifyDesc:      in.VerifyDesc,
			VerifyTime:      gtime.Now(),
			VerifyOperateId: contexts.GetUserId(ctx),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		OrderIds, _ := dao.FoodOrder.Ctx(ctx).
			Fields(dao.FoodOrder.Columns().Id).
			Where(dao.FoodOrder.Columns().SettlementOrderId, in.Id).
			Array()
		// 订单日志
		for _, OrderId := range OrderIds {
			if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
				OrderId:     g.NewVar(OrderId).Int(),
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
