package admin

import (
	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
)

func (c *ControllerSpa) OrderList(ctx context.Context, req *spa.OrderListReq) (res *spa.OrderListRes, err error) {
	list, totalCount, err := service.SpaOrder().List(ctx, &req.SpaOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaOrderListModel{}
	}

	res = new(spa.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) OrderView(ctx context.Context, req *spa.OrderViewReq) (res *spa.OrderViewRes, err error) {
	data, err := service.SpaOrder().View(ctx, &req.SpaOrderViewInp)
	if err != nil {
		return
	}

	res = new(spa.OrderViewRes)
	res.SpaOrderViewModel = data
	return
}
func (c *ControllerSpa) OrderConfirmAgree(ctx context.Context, req *spa.OrderConfirmAgreeReq) (res *spa.OrderConfirmAgreeRes, err error) {
	err = service.SpaOrder().ConfirmAgree(ctx, &req.SpaOrderConfirmAgreeInp)
	return
}
func (c *ControllerSpa) OrderConfirmDisagree(ctx context.Context, req *spa.OrderConfirmDisagreeReq) (res *spa.OrderConfirmDisagreeRes, err error) {
	err = service.SpaOrder().ConfirmDisagree(ctx, &req.SpaOrderConfirmDisagreeInp)
	return
}
func (c *ControllerSpa) OrderTechnician(ctx context.Context, req *spa.OrderTechnicianReq) (res *spa.OrderTechnicianRes, err error) {
	list, totalCount, err := service.SpaOrder().TechnicianList(ctx, &req.SpaOrderTechnicianInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaOrderTechnicianModel{}
	}

	res = new(spa.OrderTechnicianRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) OrderDispatch(ctx context.Context, req *spa.OrderDispatchReq) (res *spa.OrderDispatchRes, err error) {
	err = service.SpaOrder().Dispatch(ctx, &req.SpaOrderDispatchInp)
	return
}

// TechnicianOrderList 技师结算订单列表
func (c *ControllerSpa) TechnicianOrderList(ctx context.Context, req *spa.TechnicianOrderListReq) (res *spa.TechnicianOrderListRes, err error) {
	list, totalCount, err := service.SpaOrder().TechnicianOrderList(ctx, &req.SettleSpaOrderTechnicianListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SettleSpaOrderTechnicianListModel{}
	}

	res = new(spa.TechnicianOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) OrderGoOut(ctx context.Context, req *spa.OrderGoOutReq) (res *spa.OrderGoOutRes, err error) {
	err = service.SpaOrder().GoOut(ctx, &req.SpaOrderGoOutInp)
	return
}
func (c *ControllerSpa) OrderStartService(ctx context.Context, req *spa.OrderStartServiceReq) (res *spa.OrderStartServiceRes, err error) {
	err = service.SpaOrder().StartService(ctx, &req.SpaOrderStartServiceInp)
	return
}
func (c *ControllerSpa) OrderEndService(ctx context.Context, req *spa.OrderEndServiceReq) (res *spa.OrderEndServiceRes, err error) {
	err = service.SpaOrder().EndService(ctx, &req.SpaOrderEndServiceInp)
	return
}
func (c *ControllerSpa) OrderCancelPay(ctx context.Context, req *spa.OrderCancelPayReq) (res *spa.OrderCancelPayRes, err error) {
	err = service.SpaOrder().CancelPay(ctx, &req.SpaOrderCancelPayInp)
	return
}
func (c *ControllerSpa) OrderExport(ctx context.Context, req *spa.OrderExportReq) (res *spa.OrderExportRes, err error) {
	err = service.SpaOrder().ExportOrder(ctx, &req.SpaOrderExportInp)
	return
}
func (c *ControllerSpa) OrderExportList(ctx context.Context, req *spa.OrderExportListReq) (res *spa.OrderExportListRes, err error) {
	list, totalCount, err := service.SpaOrder().ExportList(ctx, &req.SpaOrderExportListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaOrderExportListModel{}
	}

	res = new(spa.OrderExportListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) OrderAbnormal(ctx context.Context, req *spa.OrderAbnormalReq) (res *spa.OrderAbnormalRes, err error) {
	err = service.SpaOrder().Abnormal(ctx, &req.SpaOrderAbnormalInp)
	return
}
func (c *ControllerSpa) OrderArrive(ctx context.Context, req *spa.OrderArriveReq) (res *spa.OrderArriveRes, err error) {
	err = service.SpaOrder().Arrive(ctx, &req.SpaOrderGoOutInp)
	return
}
