package app

import (
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/service"
	"context"

	"APT/api/app/cabinet"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

func (c *ControllerCabinet) OrderList(ctx context.Context, req *cabinet.OrderListReq) (res *cabinet.OrderListRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(cabinet.OrderListRes)
	if res.List, res.Count, err = service.CabinetService().AppList(ctx, &input_cabinet.OrderAppListInp{
		PageReq:     req.PageReq,
		OrderStatus: req.OrderStatus,
		MemberId:    MemberInfo.Id,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) OrderView(ctx context.Context, req *cabinet.OrderViewReq) (res *cabinet.OrderViewRes, err error) {
	res = new(cabinet.OrderViewRes)
	if res.OrderAppViewModel, err = service.CabinetService().AppView(ctx, &req.OrderAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) PayOvertime(ctx context.Context, req *cabinet.PayOvertimeReq) (res *cabinet.PayOvertimeRes, err error) {
	res = new(cabinet.PayOvertimeRes)
	if res.PayOvertimeModel, err = service.CabinetService().PayOvertime(ctx, req.PayOvertimeInp); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) CancelOrderNoPay(ctx context.Context, req *cabinet.CancelOrderNoPayReq) (res *cabinet.CancelOrderNoPayRes, err error) {
	if err = service.CabinetService().OrderExpiration(ctx, req.OrderSn); err != nil {
		// 订单过期处理失败
		err = gerror.New(gi18n.T(ctx, "order_expire_handle_failed"))
		return
	}
	return
}
func (c *ControllerCabinet) PayOvertimeInfo(ctx context.Context, req *cabinet.PayOvertimeInfoReq) (res *cabinet.PayOvertimeInfoRes, err error) {
	res = new(cabinet.PayOvertimeInfoRes)
	if res.PayOvertimeInfoModel, err = service.CabinetService().PayOvertimeInfo(ctx, &req.PayOvertimeInfoInp); err != nil {
		return
	}
	return
}
