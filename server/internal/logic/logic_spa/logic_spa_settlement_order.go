package logic_spa

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_spa"
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

type sSpaSettlementOrder struct{}

func NewSpaSettlementOrder() *sSpaSettlementOrder {
	return &sSpaSettlementOrder{}
}

func init() {
	service.RegisterSpaSettlementOrder(NewSpaSettlementOrder())
}

// Model 结算模式ORM模型
func (s *sSpaSettlementOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaSettlementOrder.Ctx(ctx), option...)
}

// DailySettlement 每日结算
func (s *sSpaSettlementOrder) DailySettlement(ctx context.Context) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartTime *gtime.Time
			EndTime   *gtime.Time
		)

		StartTime = gtime.New(gtime.Now().StartOfDay()).Add(time.Duration(-24) * time.Hour)
		EndTime = gtime.Now().StartOfDay()

		// 获取技师结算信息(周期性按日结算)
		var SettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().TechnicianId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 1).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "TECHNICIAN").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().TechnicianId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		// 服务商结算
		var IspSettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().IspId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 1).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "ISP").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().IspId).
			Scan(&IspSettlementList)
		if err != nil {
			return err
		}

		SettlementList = append(SettlementList, IspSettlementList...)

		// 遍历结算信息
		for _, SettlementInfo := range SettlementList {
			var SettlementObject string
			var IspId = gvar.New(SettlementInfo.IspId).Int()

			// 结算对象-服务商
			SettlementObject = "ISP"
			// 技师
			if SettlementInfo.TechnicianId > 0 {
				// 找到该技师信息所属服务商
				var TechnicianInfo *entity.SpaTechnician

				if err = g.Model(dao.SpaTechnician.Table()).Ctx(ctx).Safe().Scan(&TechnicianInfo); err != nil {
					err = gerror.Wrap(err, "获取技师信息失败，请稍后重试！")
					return
				}

				// 结算对象-技师
				SettlementObject = "TECHNICIAN"
				IspId = gvar.New(TechnicianInfo.IspId).Int()

			}
			InsertSettlementMap := g.MapStrAny{
				dao.SpaSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("SS"),
				dao.SpaSettlementOrder.Columns().TechnicianId:     gvar.New(SettlementInfo.TechnicianId).Int(),
				dao.SpaSettlementOrder.Columns().IspId:            IspId,
				dao.SpaSettlementOrder.Columns().SettlementObject: SettlementObject,
				dao.SpaSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.SpaSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.SpaSettlementOrder.Columns().Status:           "DONE",
				dao.SpaSettlementOrder.Columns().StartTime:        StartTime,
				dao.SpaSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_spa.SpaSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 获取关联订单ids
			settleOrderIds, _ := service.SpaOrderTechnician().GetOrderIds(ctx, SettlementInfo.TechnicianId, SettlementInfo.IspId, 2, 1, SettlementObject, gvar.New(StartTime).String(), gvar.New(EndTime).String())

			// 修改订单结算状态
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).
				WhereIn(dao.SpaOrderTechnician.Columns().OrderId, settleOrderIds).
				Data(g.MapStrAny{
					dao.SpaOrderTechnician.Columns().SettlementStatus:  "SUCCESS",
					dao.SpaOrderTechnician.Columns().SettlementTime:    gtime.Now(),
					dao.SpaOrderTechnician.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单技师表结算状态失败，请稍后重试！")
			}

			if SettlementObject == "ISP" {

				// 更新服务商预定量和预定金额
				if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, SettlementInfo.IspId).Update(g.MapStrAny{
					dao.SpaIsp.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaIsp.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaIsp.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}

			} else {
				// 更新技师预定量和预定金额
				if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, SettlementInfo.TechnicianId).Update(g.MapStrAny{
					dao.SpaTechnician.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaTechnician.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaTechnician.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}
			}

			// 订单日志
			for _, OrderId := range settleOrderIds {
				if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
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
func (s *sSpaSettlementOrder) WeekSettlement(ctx context.Context) (err error) {
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

		// 获取结算信息-技师
		var SettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().TechnicianId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 2).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "TECHNICIAN").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().TechnicianId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		// 服务商结算
		var IspSettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().IspId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 2).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "ISP").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().IspId).
			Scan(&IspSettlementList)
		if err != nil {
			return err
		}

		SettlementList = append(SettlementList, IspSettlementList...)

		for _, SettlementInfo := range SettlementList {

			var SettlementObject string
			var IspId = gvar.New(SettlementInfo.IspId).Int()

			// 结算对象-服务商
			SettlementObject = "ISP"

			// 技师
			if SettlementInfo.TechnicianId > 0 {
				// 找到该技师信息所属服务商
				var TechnicianInfo *entity.SpaTechnician

				if err = g.Model(dao.SpaTechnician.Table()).Ctx(ctx).Safe().Scan(&TechnicianInfo); err != nil {
					err = gerror.Wrap(err, "获取技师信息失败，请稍后重试！")
					return
				}

				// 结算对象-技师
				SettlementObject = "TECHNICIAN"
				IspId = gvar.New(TechnicianInfo.IspId).Int()

			}

			InsertSettlementMap := g.MapStrAny{
				dao.SpaSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("SS"),
				dao.SpaSettlementOrder.Columns().TechnicianId:     gvar.New(SettlementInfo.TechnicianId).Int(),
				dao.SpaSettlementOrder.Columns().IspId:            IspId,
				dao.SpaSettlementOrder.Columns().SettlementObject: SettlementObject,
				dao.SpaSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.SpaSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.SpaSettlementOrder.Columns().Status:           "DONE",
				dao.SpaSettlementOrder.Columns().StartTime:        StartTime,
				dao.SpaSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_spa.SpaSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 获取关联订单ids
			settleOrderIds, _ := service.SpaOrderTechnician().GetOrderIds(ctx, SettlementInfo.TechnicianId, SettlementInfo.IspId, 2, 2, SettlementObject, gvar.New(StartTime).String(), gvar.New(EndTime).String())

			// 修改订单结算状态
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).
				WhereIn(dao.SpaOrderTechnician.Columns().OrderId, settleOrderIds).
				Data(g.MapStrAny{
					dao.SpaOrderTechnician.Columns().SettlementStatus:  "SUCCESS",
					dao.SpaOrderTechnician.Columns().SettlementTime:    gtime.Now(),
					dao.SpaOrderTechnician.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单技师表结算状态失败，请稍后重试！")
			}

			if SettlementObject == "ISP" {

				// 更新服务商预定量和预定金额
				if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, SettlementInfo.IspId).Update(g.MapStrAny{
					dao.SpaIsp.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaIsp.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaIsp.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}

			} else {
				// 更新技师预定量和预定金额
				if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, SettlementInfo.TechnicianId).Update(g.MapStrAny{
					dao.SpaTechnician.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaTechnician.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaTechnician.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}
			}

			// 订单日志
			for _, OrderId := range settleOrderIds {
				if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
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
func (s *sSpaSettlementOrder) MonthSettlement(ctx context.Context) (err error) {
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

		// 获取结算信息-技师
		var SettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().TechnicianId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 3).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "TECHNICIAN").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().TechnicianId).
			Scan(&SettlementList)
		if err != nil {
			return err
		}

		// 服务商结算
		var IspSettlementList []*input_spa.SpaOrderSettleInfoModel
		err = dao.SpaOrderTechnician.Ctx(ctx).
			Fields(dao.SpaOrderTechnician.Columns().IspId).
			Fields(gdb.Raw("SUM( settlement_amount ) as totalSettlementAmount")).
			Fields(gdb.Raw("SUM( order_amount ) as totalOrderAmount")).
			Fields(gdb.Raw("COUNT( order_id ) as totalOrderNum")).
			Where(dao.SpaOrderTechnician.Columns().SettlementType, 2).
			Where(dao.SpaOrderTechnician.Columns().SettlementCycle, 3).
			WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
			WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
			WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
			Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
			Where(dao.SpaOrderTechnician.Columns().SettlementObject, "ISP").
			OmitEmptyWhere().
			Group(dao.SpaOrderTechnician.Columns().IspId).
			Scan(&IspSettlementList)
		if err != nil {
			return err
		}

		SettlementList = append(SettlementList, IspSettlementList...)

		for _, SettlementInfo := range SettlementList {

			var SettlementObject string
			var IspId = gvar.New(SettlementInfo.IspId).Int()

			// 结算对象-服务商
			SettlementObject = "ISP"
			// 技师
			if SettlementInfo.TechnicianId > 0 {
				// 找到该技师信息所属服务商
				var TechnicianInfo *entity.SpaTechnician

				if err = g.Model(dao.SpaTechnician.Table()).Ctx(ctx).Safe().Scan(&TechnicianInfo); err != nil {
					err = gerror.Wrap(err, "获取技师信息失败，请稍后重试！")
					return
				}

				// 结算对象-技师
				SettlementObject = "TECHNICIAN"
				IspId = gvar.New(TechnicianInfo.IspId).Int()

			}

			InsertSettlementMap := g.MapStrAny{
				dao.SpaSettlementOrder.Columns().OrderSn:          uuid.CreateOrderCode("SS"),
				dao.SpaSettlementOrder.Columns().TechnicianId:     gvar.New(SettlementInfo.TechnicianId).Int(),
				dao.SpaSettlementOrder.Columns().IspId:            IspId,
				dao.SpaSettlementOrder.Columns().SettlementObject: SettlementObject,
				dao.SpaSettlementOrder.Columns().OrderAmount:      gvar.New(SettlementInfo.TotalOrderAmount).Float64(),
				dao.SpaSettlementOrder.Columns().SettlementAmount: gvar.New(SettlementInfo.TotalSettlementAmount).Float64(),
				dao.SpaSettlementOrder.Columns().Status:           "DONE",
				dao.SpaSettlementOrder.Columns().StartTime:        StartTime,
				dao.SpaSettlementOrder.Columns().EndTime:          EndTime,
			}

			// 新增结算单
			var SettlementOrderId int64
			if SettlementOrderId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(input_spa.SpaSettlementOrderInsertFields{}).
				Data(InsertSettlementMap).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "新增结算单失败，请稍后重试！")
			}

			// 获取关联订单ids
			settleOrderIds, _ := service.SpaOrderTechnician().GetOrderIds(ctx, SettlementInfo.TechnicianId, SettlementInfo.IspId, 2, 3, SettlementObject, gvar.New(StartTime).String(), gvar.New(EndTime).String())

			// 修改订单结算状态
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).
				WhereIn(dao.SpaOrderTechnician.Columns().OrderId, settleOrderIds).
				Data(g.MapStrAny{
					dao.SpaOrderTechnician.Columns().SettlementStatus:  "SUCCESS",
					dao.SpaOrderTechnician.Columns().SettlementTime:    gtime.Now(),
					dao.SpaOrderTechnician.Columns().SettlementOrderId: SettlementOrderId,
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改订单技师表结算状态失败，请稍后重试！")
			}

			if SettlementObject == "ISP" {
				// 更新服务商预定量和预定金额
				if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, SettlementInfo.IspId).Update(g.MapStrAny{
					dao.SpaIsp.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaIsp.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaIsp.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}

			} else {
				// 更新技师预定量和预定金额
				if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, SettlementInfo.TechnicianId).Update(g.MapStrAny{
					dao.SpaTechnician.Columns().SettlementOrderNum:    gdb.Raw(fmt.Sprintf("settlement_order_num+%d", SettlementInfo.TotalOrderNum)),
					dao.SpaTechnician.Columns().SettlementOrderAmount: gdb.Raw(fmt.Sprintf("settlement_order_amount+%f", SettlementInfo.TotalOrderAmount)),
					dao.SpaTechnician.Columns().TotalSettlementAmount: gdb.Raw(fmt.Sprintf("total_settlement_amount+%f", SettlementInfo.TotalSettlementAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
					return
				}
			}

			// 订单日志
			for _, OrderId := range settleOrderIds {
				if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
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
func (s *sSpaSettlementOrder) SettlementOrderStat(ctx context.Context, in *input_spa.SpaSettlementOrderStatInp) (res *input_spa.SpaSettlementOrderStatModel, err error) {
	res = &input_spa.SpaSettlementOrderStatModel{}
	mod := dao.SpaOrder.Ctx(ctx)

	// 订单总额(已支付)
	res.TotalOrderAmount, err = mod.WhereNotNull(dao.SpaOrder.Columns().PayTime).Sum(dao.SpaOrder.Columns().OrderAmount)

	// 待结算金额(已支付已完成但还未结算的金额)
	res.WaitSettlementAmount, err = dao.SpaOrderTechnician.Ctx(ctx).
		Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").
		WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
		Sum(dao.SpaOrderTechnician.Columns().SettlementAmount)

	// 已结算
	res.DoneSettlementAmount, err = dao.SpaOrderTechnician.Ctx(ctx).
		Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "SUCCESS").
		Sum(dao.SpaOrder.Columns().SettlementAmount)

	// 已核账
	res.HadVerifySettlementAmount, err = dao.SpaSettlementOrder.Ctx(ctx).
		Where(dao.SpaSettlementOrder.Columns().VerifyStatus, "VERIFIED").
		Sum(dao.SpaSettlementOrder.Columns().SettlementAmount)

	return
}

// List 结算列表
func (s *sSpaSettlementOrder) List(ctx context.Context, in *input_spa.SpaSettlementOrderListInp) (list []*input_spa.SpaSettlementOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_spa.SpaSettlementOrderListModel{})

	if in.TechnicianId > 0 {
		mod = mod.Where(dao.SpaSettlementOrder.Columns().TechnicianId, in.TechnicianId)
	}

	if in.IspId > 0 {
		mod = mod.Where(dao.SpaSettlementOrder.Columns().IspId, in.IspId)
	}

	if !g.IsEmpty(in.SettlementObject) {
		mod = mod.Where(dao.SpaSettlementOrder.Columns().SettlementObject, in.SettlementObject)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.SpaSettlementOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.VerifyStatus) {
		mod = mod.Where(dao.SpaSettlementOrder.Columns().VerifyStatus, in.VerifyStatus)
	}

	if len(in.VerifyTime) == 2 {
		mod = mod.WhereBetween(dao.SpaSettlementOrder.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.SpaSettlementOrder.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算单列表失败，请稍后重试！")
		return
	}
	return
}

// View 结算单详情
func (s *sSpaSettlementOrder) View(ctx context.Context, in *input_spa.SpaSettlementOrderViewInp) (res *input_spa.SpaSettlementOrderViewModel, err error) {
	if err = s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算单详情失败，请稍后重试！")
		return
	}
	return
}

// Verify 核账
func (s *sSpaSettlementOrder) Verify(ctx context.Context, in *input_spa.SpaSettlementOrderVerifyInp) (err error) {

	var models *entity.SpaSettlementOrder
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
			WherePri(in.Id).Data(input_spa.SpaSettlementOrderVerifyFields{
			VerifyStatus:    "VERIFIED",
			VerifyImg:       in.VerifyImg,
			VerifyDesc:      in.VerifyDesc,
			VerifyTime:      gtime.Now(),
			VerifyOperateId: contexts.GetUserId(ctx),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if models.SettlementObject == "ISP" {
			// 更新服务商预定量和预定金额
			if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, models.IspId).Update(g.MapStrAny{
				dao.SpaIsp.Columns().VerifyMoney: gdb.Raw(fmt.Sprintf("verify_money+%f", models.SettlementAmount)),
				dao.SpaIsp.Columns().Balance:     gdb.Raw(fmt.Sprintf("balance+%f", models.SettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新服务商信息失败，请稍后重试！")
				return
			}

			// 写入服务商积分变更日志
			if _, err = dao.SpaIspBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.SpaIspBalanceChange{
				IspId:             models.IspId,
				Type:              "SETTLEMENT",
				ChangePrice:       models.SettlementAmount,
				SettlementOrderId: int(models.Id),
				Des:               "订单收入",
				CreatedAt:         gtime.Now(),
				UpdatedAt:         gtime.Now(),
			}); err != nil {
				err = gerror.Wrap(err, "写入服务商积分变更日志失败，请稍后重试！")
				return
			}
		} else {
			// 更新技师预定量和预定金额
			if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, models.TechnicianId).Update(g.MapStrAny{
				dao.SpaTechnician.Columns().VerifyMoney: gdb.Raw(fmt.Sprintf("verify_money+%f", models.SettlementAmount)),
				dao.SpaTechnician.Columns().Balance:     gdb.Raw(fmt.Sprintf("balance+%f", models.SettlementAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
				return
			}

			// 写入技师积分变更日志
			if _, err = dao.SpaTechnicianBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.SpaTechnicianBalanceChange{
				TechnicianId:      models.TechnicianId,
				Type:              "SETTLEMENT",
				ChangePrice:       models.SettlementAmount,
				SettlementOrderId: int(models.Id),
				Des:               "订单收入",
				CreatedAt:         gtime.Now(),
				UpdatedAt:         gtime.Now(),
			}); err != nil {
				err = gerror.Wrap(err, "写入技师积分变更日志失败，请稍后重试！")
				return
			}
		}

		OrderIds, _ := dao.SpaOrder.Ctx(ctx).
			Fields(dao.SpaOrder.Columns().Id).
			Where(dao.SpaOrder.Columns().SettlementOrderId, in.Id).
			Array()

		// 订单日志
		for _, OrderId := range OrderIds {
			if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
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
