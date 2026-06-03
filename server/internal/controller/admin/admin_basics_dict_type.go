package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) DictTypeTree(ctx context.Context, _ *basics.DictTypeTreeReq) (res *basics.DictTypeTreeRes, err error) {
	res = new(basics.DictTypeTreeRes)
	res.List, err = service.BasicsDictType().Tree(ctx)
	return
}

func (c *ControllerBasics) DictTypeEdit(ctx context.Context, req *basics.DictTypeEditReq) (res *basics.DictTypeEditRes, err error) {
	err = service.BasicsDictType().Edit(ctx, &req.DictTypeEditInp)
	return
}

func (c *ControllerBasics) DictTypeDelete(ctx context.Context, req *basics.DictTypeDeleteReq) (res *basics.DictTypeDeleteRes, err error) {
	err = service.BasicsDictType().Delete(ctx, &req.DictTypeDeleteInp)
	return
}
