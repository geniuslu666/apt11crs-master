package admin

import (
	"APT/api/admin/cabinet"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/service"
	"context"
)

func (c *ControllerCabinet) OrderList(ctx context.Context, req *cabinet.OrderListReq) (res *cabinet.OrderListRes, err error) {
	list, totalCount, err := service.CabinetService().List(ctx, &req.OrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_cabinet.OrderListModel{}
	}

	res = new(cabinet.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCabinet) OrderView(ctx context.Context, req *cabinet.OrderViewReq) (res *cabinet.OrderViewRes, err error) {
	data, err := service.CabinetService().OrderView(ctx, &req.OrderViewInp)
	if err != nil {
		return
	}

	res = new(cabinet.OrderViewRes)
	res.OrderViewModel = data
	return
}
func (c *ControllerCabinet) OrderExport(ctx context.Context, req *cabinet.OrderExportReq) (res *cabinet.OrderExportRes, err error) {
	err = service.CabinetService().ExportOrder(ctx, &req.OrderExportInp)
	return
}
func (c *ControllerCabinet) OrderExportList(ctx context.Context, req *cabinet.OrderExportListReq) (res *cabinet.OrderExportListRes, err error) {
	list, totalCount, err := service.CabinetService().ExportList(ctx, &req.OrderExportListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_cabinet.OrderExportListModel{}
	}

	res = new(cabinet.OrderExportListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCabinet) OrderComplete(ctx context.Context, req *cabinet.OrderCompleteReq) (res *cabinet.OrderCompleteRes, err error) {
	err = service.CabinetService().CompleteOrder(ctx, &req.OrderCompleteInp)
	return
}
func (c *ControllerCabinet) OrderRefund(ctx context.Context, req *cabinet.OrderRefundReq) (res *cabinet.OrderRefundRes, err error) {
	err = service.CabinetService().Refund(ctx, &req.CabinetOrderRefundInp)
	return
}
