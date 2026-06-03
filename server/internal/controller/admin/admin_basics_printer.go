package admin

import (
	"APT/api/admin/basics"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) PrinterList(ctx context.Context, req *basics.PrinterListReq) (res *basics.PrinterListRes, err error) {
	list, totalCount, err := service.BasicsPrinter().List(ctx, &req.PrinterListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PrinterListModel{}
	}

	res = new(basics.PrinterListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)

	//err = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{})
	//if err != nil {
	//	return nil, err
	//}
	return
}
func (c *ControllerBasics) PrinterView(ctx context.Context, req *basics.PrinterViewReq) (res *basics.PrinterViewRes, err error) {
	data, err := service.BasicsPrinter().View(ctx, &req.PrinterViewInp)
	if err != nil {
		return
	}

	res = new(basics.PrinterViewRes)
	res.PrinterViewModel = data
	return
}
func (c *ControllerBasics) PrinterEdit(ctx context.Context, req *basics.PrinterEditReq) (res *basics.PrinterEditRes, err error) {
	err = service.BasicsPrinter().Edit(ctx, &req.PrinterEditInp)
	return
}
func (c *ControllerBasics) PrinterDelete(ctx context.Context, req *basics.PrinterDeleteReq) (res *basics.PrinterDeleteRes, err error) {
	err = service.BasicsPrinter().Delete(ctx, &req.PrinterDeleteInp)
	return
}
func (c *ControllerBasics) PrinterMaxSort(ctx context.Context, req *basics.PrinterMaxSortReq) (res *basics.PrinterMaxSortRes, err error) {
	data, err := service.BasicsPrinter().MaxSort(ctx, &req.PrinterMaxSortInp)
	if err != nil {
		return
	}

	res = new(basics.PrinterMaxSortRes)
	res.PrinterMaxSortModel = data
	return
}
func (c *ControllerBasics) PrinterCarOrder(ctx context.Context, req *basics.PrinterCarOrderReq) (res *basics.PrinterCarOrderRes, err error) {
	err = service.BasicsPrinter().PrinterCarOrder(ctx, &req.PrinterCarOrderInp)
	return
}
func (c *ControllerBasics) PrinterStatus(ctx context.Context, req *basics.PrinterStatusReq) (res *basics.PrinterStatusRes, err error) {
	err = service.BasicsPrinter().Status(ctx, &req.PrinterStatusInp)
	return
}
func (c *ControllerBasics) PrinterFoodOrder(ctx context.Context, req *basics.PrinterFoodOrderReq) (res *basics.PrinterFoodOrderRes, err error) {
	err = service.BasicsPrinter().PrinterFoodOrder(ctx, &req.PrinterCarOrderInp)
	return
}
func (c *ControllerBasics) PrinterSpaOrder(ctx context.Context, req *basics.PrinterSpaOrderReq) (res *basics.PrinterSpaOrderRes, err error) {
	err = service.BasicsPrinter().PrinterSpaOrder(ctx, &req.PrinterCarOrderInp)
	return
}
