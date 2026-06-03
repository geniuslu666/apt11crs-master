package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaOrderTechnician struct{}

func NewSpaOrderTechnician() *sSpaOrderTechnician {
	return &sSpaOrderTechnician{}
}

func init() {
	service.RegisterSpaOrderTechnician(NewSpaOrderTechnician())
}

func (s *sSpaOrderTechnician) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaOrderTechnician.Ctx(ctx), option...)
}

// GetOrderIds 根据技师ID、结算类型、结算周期、开始时间、结束时间获取订单ID
func (s *sSpaOrderTechnician) GetOrderIds(ctx context.Context, technicianId int, ispId int, settlementType int, settlementCycle int, settlementObject string, StartTime, EndTime string) (ids []int, err error) {

	mod := dao.SpaOrderTechnician.Ctx(ctx).
		Fields("order_id")

	// 服务商
	if settlementObject == "ISP" {
		mod = mod.Where(dao.SpaOrderTechnician.Columns().IspId, ispId)
	} else {
		mod = mod.Where(dao.SpaOrderTechnician.Columns().TechnicianId, technicianId)
	}

	columns, err := mod.
		Where(dao.SpaOrderTechnician.Columns().SettlementType, settlementType).
		Where(dao.SpaOrderTechnician.Columns().SettlementCycle, settlementCycle).
		WhereGTE(dao.SpaOrderTechnician.Columns().SettlementAmount, 0).
		WhereNotNull(dao.SpaOrderTechnician.Columns().ActualEndTime).
		WhereBetween(dao.SpaOrderTechnician.Columns().ActualEndTime, StartTime, EndTime).
		Where(dao.SpaOrderTechnician.Columns().SettlementObject, settlementObject).
		Where(dao.SpaOrderTechnician.Columns().SettlementStatus, "WAIT").Array()

	if err != nil {
		err = gerror.Wrap(err, "获取订单ID失败！")
		return
	}
	ids = g.NewVar(columns).Ints()

	return
}

// GetTechnicianIds 根据订单ID获取技师ID
func (s *sSpaOrderTechnician) GetTechnicianIds(ctx context.Context, orderId int) (ids []int, err error) {

	columns, err := dao.SpaOrderTechnician.Ctx(ctx).
		Fields("technician_id").
		Where(dao.SpaOrderTechnician.Columns().OrderId, orderId).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取技师ID失败！")
		return
	}
	ids = g.NewVar(columns).Ints()

	return
}
