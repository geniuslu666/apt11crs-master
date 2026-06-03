package admin

import (
	"APT/api/admin/th"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
)

func (c *ControllerTh) MchCategoryList(ctx context.Context, req *th.MchCategoryListReq) (res *th.MchCategoryListRes, err error) {
	list, totalCount, err := service.ThMchCategory().List(ctx, &req.ThMchCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchCategoryListModel{}
	}

	res = new(th.MchCategoryListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) MchCategoryAll(ctx context.Context, req *th.MchCategoryAllReq) (res *th.MchCategoryAllRes, err error) {
	list, err := service.ThMchCategory().All(ctx, &req.ThMchCategoryAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThMchCategoryAllModel{}
	}

	res = new(th.MchCategoryAllRes)
	res.List = list
	return
}
func (c *ControllerTh) MchCategoryView(ctx context.Context, req *th.MchCategoryViewReq) (res *th.MchCategoryViewRes, err error) {
	data, err := service.ThMchCategory().View(ctx, &req.ThMchCategoryViewInp)
	if err != nil {
		return
	}

	res = new(th.MchCategoryViewRes)
	res.ThMchCategoryViewModel = data
	return
}
func (c *ControllerTh) MchCategoryEdit(ctx context.Context, req *th.MchCategoryEditReq) (res *th.MchCategoryEditRes, err error) {
	err = service.ThMchCategory().Edit(ctx, &req.ThMchCategoryEditInp)
	return
}
func (c *ControllerTh) MchCategoryDelete(ctx context.Context, req *th.MchCategoryDeleteReq) (res *th.MchCategoryDeleteRes, err error) {
	err = service.ThMchCategory().Delete(ctx, &req.ThMchCategoryDeleteInp)
	return
}
func (c *ControllerTh) MchCategorySwitch(ctx context.Context, req *th.MchCategorySwitchReq) (res *th.MchCategorySwitchRes, err error) {
	err = service.ThMchCategory().Switch(ctx, &req.ThMchCategorySwitchInp)
	return
}
