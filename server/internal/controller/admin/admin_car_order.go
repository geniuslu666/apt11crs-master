package admin

import (
	"APT/api/admin/car"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
)

func (c *ControllerCar) OrderList(ctx context.Context, req *car.OrderListReq) (res *car.OrderListRes, err error) {
	list, totalCount, err := service.CarOrder().List(ctx, &req.CarOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarOrderListModel{}
	}

	res = new(car.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) OrderView(ctx context.Context, req *car.OrderViewReq) (res *car.OrderViewRes, err error) {
	data, err := service.CarOrder().View(ctx, &req.CarOrderViewInp)
	if err != nil {
		return
	}

	res = new(car.OrderViewRes)
	res.CarOrderViewModel = data
	return
}
func (c *ControllerCar) OrderConfirmAgree(ctx context.Context, req *car.OrderConfirmAgreeReq) (res *car.OrderConfirmAgreeRes, err error) {
	err = service.CarOrder().ConfirmAgree(ctx, &req.CarOrderConfirmAgreeInp)
	return
}
func (c *ControllerCar) OrderConfirmDisagree(ctx context.Context, req *car.OrderConfirmDisagreeReq) (res *car.OrderConfirmDisagreeRes, err error) {
	err = service.CarOrder().ConfirmDisagree(ctx, &req.CarOrderConfirmDisagreeInp)
	return
}
func (c *ControllerCar) OrderDriver(ctx context.Context, req *car.OrderDriverReq) (res *car.OrderDriverRes, err error) {
	list, totalCount, err := service.CarOrder().DriverList(ctx, &req.CarOrderDriverInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarOrderDriverModel{}
	}

	res = new(car.OrderDriverRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) OrderDispatch(ctx context.Context, req *car.OrderDispatchReq) (res *car.OrderDispatchRes, err error) {
	err = service.CarOrder().Dispatch(ctx, &req.CarOrderDispatchInp)
	return
}

// SettleOrderList 结算订单列表
func (c *ControllerCar) SettleOrderList(ctx context.Context, req *car.SettleOrderListReq) (res *car.SettleOrderListRes, err error) {
	list, totalCount, err := service.CarOrder().SettleOrderList(ctx, &req.SettleCarOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.SettleCarOrderListModel{}
	}

	res = new(car.SettleOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) OrderRefund(ctx context.Context, req *car.OrderRefundReq) (res *car.OrderRefundRes, err error) {
	err = service.CarOrder().Refund(ctx, &req.CarOrderRefundInp)

	return
}
func (c *ControllerCar) OrderStartService(ctx context.Context, req *car.OrderStartServiceReq) (res *car.OrderStartServiceRes, err error) {
	err = service.CarOrder().StartService(ctx, &req.CarOrderStartServiceInp)
	return
}
func (c *ControllerCar) OrderEndService(ctx context.Context, req *car.OrderEndServiceReq) (res *car.OrderEndServiceRes, err error) {
	err = service.CarOrder().EndService(ctx, &req.CarOrderEndServiceInp)
	return
}
func (c *ControllerCar) OrderGoOut(ctx context.Context, req *car.OrderGoOutReq) (res *car.OrderGoOutRes, err error) {
	err = service.CarOrder().GoOut(ctx, &req.CarOrderGoOutInp)
	return
}
func (c *ControllerCar) OrderAbnormal(ctx context.Context, req *car.OrderAbnormalReq) (res *car.OrderAbnormalRes, err error) {
	err = service.CarOrder().Abnormal(ctx, &req.CarOrderAbnormalInp)
	return
}
func (c *ControllerCar) OrderExport(ctx context.Context, req *car.OrderExportReq) (res *car.OrderExportRes, err error) {
	err = service.CarOrder().ExportOrder(ctx, &req.CarOrderExportInp)
	return
}
func (c *ControllerCar) OrderExportList(ctx context.Context, req *car.OrderExportListReq) (res *car.OrderExportListRes, err error) {
	list, totalCount, err := service.CarOrder().ExportList(ctx, &req.CarOrderExportListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarOrderExportListModel{}
	}

	res = new(car.OrderExportListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) OrderTransfer(ctx context.Context, req *car.OrderTransferReq) (res *car.OrderTransferRes, err error) {
	err = service.CarOrder().TransferOrder(ctx, &req.CarOrderTransferInp)
	return
}
