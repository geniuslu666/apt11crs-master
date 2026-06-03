package admin

import (
	"APT/api/admin/basics"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) IndexBannerList(ctx context.Context, req *basics.IndexBannerListReq) (res *basics.IndexBannerListRes, err error) {
	list, totalCount, err := service.BasicsIndexBanner().List(ctx, &req.IndexBannerListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.IndexBannerListModel{}
	}

	res = new(basics.IndexBannerListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) IndexBannerView(ctx context.Context, req *basics.IndexBannerViewReq) (res *basics.IndexBannerViewRes, err error) {
	data, err := service.BasicsIndexBanner().View(ctx, &req.IndexBannerViewInp)
	if err != nil {
		return
	}

	res = new(basics.IndexBannerViewRes)
	res.IndexBannerViewModel = data
	return
}
func (c *ControllerBasics) IndexBannerEdit(ctx context.Context, req *basics.IndexBannerEditReq) (res *basics.IndexBannerEditRes, err error) {
	err = service.BasicsIndexBanner().Edit(ctx, &req.IndexBannerEditInp)
	return
}
func (c *ControllerBasics) IndexBannerDelete(ctx context.Context, req *basics.IndexBannerDeleteReq) (res *basics.IndexBannerDeleteRes, err error) {
	err = service.BasicsIndexBanner().Delete(ctx, &req.IndexBannerDeleteInp)
	return
}
func (c *ControllerBasics) IndexBannerMaxSort(ctx context.Context, req *basics.IndexBannerMaxSortReq) (res *basics.IndexBannerMaxSortRes, err error) {
	data, err := service.BasicsIndexBanner().MaxSort(ctx, &req.IndexBannerMaxSortInp)
	if err != nil {
		return
	}

	res = new(basics.IndexBannerMaxSortRes)
	res.IndexBannerMaxSortModel = data
	return
}
func (c *ControllerBasics) IndexBannerStatus(ctx context.Context, req *basics.IndexBannerStatusReq) (res *basics.IndexBannerStatusRes, err error) {
	err = service.BasicsIndexBanner().Status(ctx, &req.IndexBannerStatusInp)
	return
}
func (c *ControllerBasics) IndexBannerSort(ctx context.Context, req *basics.IndexBannerSortReq) (res *basics.IndexBannerSortRes, err error) {
	err = service.BasicsIndexBanner().BannerSort(ctx, &req.IndexBannerSortInp)
	return
}
