package admin

import (
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/th"
)

func (c *ControllerTh) MchStoreList(ctx context.Context, req *th.MchStoreListReq) (res *th.MchStoreListRes, err error) {
	list, totalCount, err := service.ThMchStore().List(ctx, &req.ThMchStoreListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchStoreListModel{}
	}

	res = new(th.MchStoreListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) MchStoreView(ctx context.Context, req *th.MchStoreViewReq) (res *th.MchStoreViewRes, err error) {
	data, err := service.ThMchStore().View(ctx, &req.ThMchStoreViewInp)
	if err != nil {
		return
	}

	res = new(th.MchStoreViewRes)
	res.ThMchStoreViewModel = data
	return
}
func (c *ControllerTh) MchStoreEdit(ctx context.Context, req *th.MchStoreEditReq) (res *th.MchStoreEditRes, err error) {
	err = service.ThMchStore().Edit(ctx, &req.ThMchStoreEditInp)
	return
}
func (c *ControllerTh) MchStoreDelete(ctx context.Context, req *th.MchStoreDeleteReq) (res *th.MchStoreDeleteRes, err error) {
	err = service.ThMchStore().Delete(ctx, &req.ThMchStoreDeleteInp)
	return
}
func (c *ControllerTh) MchStoreStatus(ctx context.Context, req *th.MchStoreStatusReq) (res *th.MchStoreStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerTh) MchStoreSwitch(ctx context.Context, req *th.MchStoreSwitchReq) (res *th.MchStoreSwitchRes, err error) {
	err = service.ThMchStore().Switch(ctx, &req.ThMchStoreSwitchInp)
	return
}
func (c *ControllerTh) MchStoreAllList(ctx context.Context, req *th.MchStoreAllListReq) (res *th.MchStoreAllListRes, err error) {
	list, err := service.ThMchStore().StoreAll(ctx, &req.ThMchStoreAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchStoreAllModel{}
	}

	res = new(th.MchStoreAllListRes)
	res.List = list
	return
}
