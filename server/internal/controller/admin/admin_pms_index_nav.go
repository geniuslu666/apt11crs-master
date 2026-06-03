package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
)

func (c *ControllerPms) IndexNavList(ctx context.Context, req *pms.IndexNavListReq) (res *pms.IndexNavListRes, err error) {
	list, totalCount, err := service.BasicsIndexNav().List(ctx, &req.PmsIndexNavListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsIndexNavListModel{}
	}

	res = new(pms.IndexNavListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) IndexNavAll(ctx context.Context, req *pms.IndexNavAllReq) (res *pms.IndexNavAllRes, err error) {
	list, err := service.BasicsIndexNav().All(ctx, &req.PmsIndexNavAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsIndexNavAllModel{}
	}

	res = new(pms.IndexNavAllRes)
	res.List = list
	return
}
func (c *ControllerPms) IndexNavView(ctx context.Context, req *pms.IndexNavViewReq) (res *pms.IndexNavViewRes, err error) {
	data, err := service.BasicsIndexNav().View(ctx, &req.PmsIndexNavViewInp)
	if err != nil {
		return
	}

	res = new(pms.IndexNavViewRes)
	res.PmsIndexNavViewModel = data
	return
}
func (c *ControllerPms) IndexNavEdit(ctx context.Context, req *pms.IndexNavEditReq) (res *pms.IndexNavEditRes, err error) {
	err = service.BasicsIndexNav().Edit(ctx, &req.PmsIndexNavEditInp)
	return
}
func (c *ControllerPms) IndexNavDelete(ctx context.Context, req *pms.IndexNavDeleteReq) (res *pms.IndexNavDeleteRes, err error) {
	err = service.BasicsIndexNav().Delete(ctx, &req.PmsIndexNavDeleteInp)
	return
}
func (c *ControllerPms) IndexNavStatus(ctx context.Context, req *pms.IndexNavStatusReq) (res *pms.IndexNavStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerPms) IndexNavSwitch(ctx context.Context, req *pms.IndexNavSwitchReq) (res *pms.IndexNavSwitchRes, err error) {
	err = service.BasicsIndexNav().Switch(ctx, &req.PmsIndexNavSwitchInp)
	return
}
func (c *ControllerPms) IndexNavSort(ctx context.Context, req *pms.IndexNavSortReq) (res *pms.IndexNavSortRes, err error) {
	err = service.BasicsIndexNav().NavSort(ctx, &req.PmsIndexNavSortInp)
	return
}
func (c *ControllerPms) IndexNavMinappStatus(ctx context.Context, req *pms.IndexNavMinappStatusReq) (res *pms.IndexNavMinappStatusRes, err error) {
	err = service.BasicsIndexNav().MinappStatus(ctx, &req.PmsIndexNavSwitchInp)
	return
}
