package admin

import (
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/th"
)

func (c *ControllerTh) MchList(ctx context.Context, req *th.MchListReq) (res *th.MchListRes, err error) {
	list, totalCount, err := service.ThMch().List(ctx, &req.ThMchListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchListModel{}
	}

	res = new(th.MchListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) MchView(ctx context.Context, req *th.MchViewReq) (res *th.MchViewRes, err error) {
	data, err := service.ThMch().View(ctx, &req.ThMchViewInp)
	if err != nil {
		return
	}

	res = new(th.MchViewRes)
	res.ThMchViewModel = data
	return
}
func (c *ControllerTh) MchEdit(ctx context.Context, req *th.MchEditReq) (res *th.MchEditRes, err error) {
	err = service.ThMch().Edit(ctx, &req.ThMchEditInp)
	return
}
func (c *ControllerTh) MchDelete(ctx context.Context, req *th.MchDeleteReq) (res *th.MchDeleteRes, err error) {
	err = service.ThMch().Delete(ctx, &req.ThMchDeleteInp)
	return
}
func (c *ControllerTh) MchStatus(ctx context.Context, req *th.MchStatusReq) (res *th.MchStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerTh) MchSwitch(ctx context.Context, req *th.MchSwitchReq) (res *th.MchSwitchRes, err error) {
	err = service.ThMch().Switch(ctx, &req.ThMchSwitchInp)
	return
}
func (c *ControllerTh) MchAll(ctx context.Context, req *th.MchAllReq) (res *th.MchAllRes, err error) {
	list, err := service.ThMch().All(ctx, &req.ThMchAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchAllModel{}
	}

	res = new(th.MchAllRes)
	res.List = list
	return
}
