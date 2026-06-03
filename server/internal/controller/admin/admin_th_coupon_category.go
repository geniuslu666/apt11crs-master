package admin

import (
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"

	"APT/api/admin/th"
)

func (c *ControllerTh) CouponCategoryList(ctx context.Context, req *th.CouponCategoryListReq) (res *th.CouponCategoryListRes, err error) {
	list, totalCount, err := service.ThCouponCategory().List(ctx, &req.ThCouponCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThCouponCategoryListModel{}
	}

	res = new(th.CouponCategoryListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) CouponCategoryAll(ctx context.Context, req *th.CouponCategoryAllReq) (res *th.CouponCategoryAllRes, err error) {
	list, err := service.ThCouponCategory().All(ctx, &req.ThCouponCategoryAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThCouponCategoryAllModel{}
	}

	res = new(th.CouponCategoryAllRes)
	res.List = list
	return
}
func (c *ControllerTh) CouponCategoryView(ctx context.Context, req *th.CouponCategoryViewReq) (res *th.CouponCategoryViewRes, err error) {
	data, err := service.ThCouponCategory().View(ctx, &req.ThCouponCategoryViewInp)
	if err != nil {
		return
	}

	res = new(th.CouponCategoryViewRes)
	res.ThCouponCategoryViewModel = data
	return
}
func (c *ControllerTh) CouponCategoryEdit(ctx context.Context, req *th.CouponCategoryEditReq) (res *th.CouponCategoryEditRes, err error) {
	err = service.ThCouponCategory().Edit(ctx, &req.ThCouponCategoryEditInp)
	return
}
func (c *ControllerTh) CouponCategoryDelete(ctx context.Context, req *th.CouponCategoryDeleteReq) (res *th.CouponCategoryDeleteRes, err error) {
	err = service.ThCouponCategory().Delete(ctx, &req.ThCouponCategoryDeleteInp)
	return
}
func (c *ControllerTh) CouponCategorySwitch(ctx context.Context, req *th.CouponCategorySwitchReq) (res *th.CouponCategorySwitchRes, err error) {
	err = service.ThCouponCategory().Switch(ctx, &req.ThCouponCategorySwitchInp)
	return
}
