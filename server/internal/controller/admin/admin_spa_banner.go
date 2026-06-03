package admin

import (
	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
)

func (c *ControllerSpa) BannerList(ctx context.Context, req *spa.BannerListReq) (res *spa.BannerListRes, err error) {
	list, totalCount, err := service.SpaBanner().List(ctx, &req.SpaBannerListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaBannerListModel{}
	}

	res = new(spa.BannerListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) BannerView(ctx context.Context, req *spa.BannerViewReq) (res *spa.BannerViewRes, err error) {
	data, err := service.SpaBanner().View(ctx, &req.SpaBannerViewInp)
	if err != nil {
		return
	}

	res = new(spa.BannerViewRes)
	res.SpaBannerViewModel = data
	return
}
func (c *ControllerSpa) BannerEdit(ctx context.Context, req *spa.BannerEditReq) (res *spa.BannerEditRes, err error) {
	err = service.SpaBanner().Edit(ctx, &req.SpaBannerEditInp)
	return
}
func (c *ControllerSpa) BannerDelete(ctx context.Context, req *spa.BannerDeleteReq) (res *spa.BannerDeleteRes, err error) {
	err = service.SpaBanner().Delete(ctx, &req.SpaBannerDeleteInp)
	return
}
func (c *ControllerSpa) BannerMaxSort(ctx context.Context, req *spa.BannerMaxSortReq) (res *spa.BannerMaxSortRes, err error) {
	data, err := service.SpaBanner().MaxSort(ctx, &req.SpaBannerMaxSortInp)
	if err != nil {
		return
	}

	res = new(spa.BannerMaxSortRes)
	res.SpaBannerMaxSortModel = data
	return
}
func (c *ControllerSpa) BannerStatus(ctx context.Context, req *spa.BannerStatusReq) (res *spa.BannerStatusRes, err error) {
	err = service.SpaBanner().Status(ctx, &req.SpaBannerStatusInp)
	return
}
