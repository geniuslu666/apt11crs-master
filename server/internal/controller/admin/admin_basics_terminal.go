package admin

import (
	"APT/api/admin/basics"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) TerminalList(ctx context.Context, req *basics.TerminalListReq) (res *basics.TerminalListRes, err error) {
	list, totalCount, err := service.BasicsTerminal().List(ctx, &req.TerminalListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.TerminalListModel{}
	}

	res = new(basics.TerminalListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)

	return
}
func (c *ControllerBasics) TerminalView(ctx context.Context, req *basics.TerminalViewReq) (res *basics.TerminalViewRes, err error) {
	data, err := service.BasicsTerminal().View(ctx, &req.TerminalViewInp)
	if err != nil {
		return
	}

	res = new(basics.TerminalViewRes)
	res.TerminalViewModel = data
	return
}
func (c *ControllerBasics) TerminalEdit(ctx context.Context, req *basics.TerminalEditReq) (res *basics.TerminalEditRes, err error) {
	err = service.BasicsTerminal().Edit(ctx, &req.TerminalEditInp)
	return
}
func (c *ControllerBasics) TerminalDelete(ctx context.Context, req *basics.TerminalDeleteReq) (res *basics.TerminalDeleteRes, err error) {
	err = service.BasicsTerminal().Delete(ctx, &req.TerminalDeleteInp)
	return
}
func (c *ControllerBasics) BrandList(ctx context.Context, req *basics.BrandListReq) (res *basics.BrandListRes, err error) {
	list, totalCount, err := service.BasicsTerminal().BrandList(ctx, &req.BrandListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.BrandListModel{}
	}

	res = new(basics.BrandListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)

	return
}
func (c *ControllerBasics) BrandView(ctx context.Context, req *basics.BrandViewReq) (res *basics.BrandViewRes, err error) {
	data, err := service.BasicsTerminal().BrandView(ctx, &req.BrandViewInp)
	if err != nil {
		return
	}

	res = new(basics.BrandViewRes)
	res.BrandViewModel = data
	return
}
func (c *ControllerBasics) BrandEdit(ctx context.Context, req *basics.BrandEditReq) (res *basics.BrandEditRes, err error) {
	err = service.BasicsTerminal().BrandEdit(ctx, &req.BrandEditInp)
	return
}
func (c *ControllerBasics) PrintTest(ctx context.Context, req *basics.PrintTestReq) (res *basics.PrintTestRes, err error) {
	err = service.BasicsTerminal().PrinterTest(ctx, &req.PrinterTestInp)
	return
}
