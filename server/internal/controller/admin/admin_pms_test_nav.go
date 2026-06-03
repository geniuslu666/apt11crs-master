package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
)

func (c *ControllerPms) TestNavList(ctx context.Context, req *pms.TestNavListReq) (res *pms.TestNavListRes, err error) {
	list, totalCount, err := service.BasicsTestNav().List(ctx, &req.PmsTestNavListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsTestNavListModel{}
	}

	res = new(pms.TestNavListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) TestNavAll(ctx context.Context, req *pms.TestNavAllReq) (res *pms.TestNavAllRes, err error) {
	list, err := service.BasicsTestNav().All(ctx, &req.PmsTestNavAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsTestNavAllModel{}
	}

	res = new(pms.TestNavAllRes)
	res.List = list
	return
}
func (c *ControllerPms) TestNavView(ctx context.Context, req *pms.TestNavViewReq) (res *pms.TestNavViewRes, err error) {
	data, err := service.BasicsTestNav().View(ctx, &req.PmsTestNavViewInp)
	if err != nil {
		return
	}

	res = new(pms.TestNavViewRes)
	res.PmsTestNavViewModel = data
	return
}
func (c *ControllerPms) TestNavEdit(ctx context.Context, req *pms.TestNavEditReq) (res *pms.TestNavEditRes, err error) {
	err = service.BasicsTestNav().Edit(ctx, &req.PmsTestNavEditInp)
	return
}
func (c *ControllerPms) TestNavDelete(ctx context.Context, req *pms.TestNavDeleteReq) (res *pms.TestNavDeleteRes, err error) {
	err = service.BasicsTestNav().Delete(ctx, &req.PmsTestNavDeleteInp)
	return
}
func (c *ControllerPms) TestNavStatus(ctx context.Context, req *pms.TestNavStatusReq) (res *pms.TestNavStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerPms) TestNavSwitch(ctx context.Context, req *pms.TestNavSwitchReq) (res *pms.TestNavSwitchRes, err error) {
	err = service.BasicsTestNav().Switch(ctx, &req.PmsTestNavSwitchInp)
	return
}
func (c *ControllerPms) TestNavSort(ctx context.Context, req *pms.TestNavSortReq) (res *pms.TestNavSortRes, err error) {
	err = service.BasicsTestNav().NavSort(ctx, &req.PmsTestNavSortInp)
	return
}
