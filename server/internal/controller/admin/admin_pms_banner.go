package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) BannerList(ctx context.Context, req *pms.BannerListReq) (res *pms.BannerListRes, err error) {
	list, totalCount, err := service.BasicsBanner().List(ctx, &req.PmsBannerListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsBannerListModel{}
	}

	res = new(pms.BannerListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) BannerView(ctx context.Context, req *pms.BannerViewReq) (res *pms.BannerViewRes, err error) {
	data, err := service.BasicsBanner().View(ctx, &req.PmsBannerViewInp)
	if err != nil {
		return
	}

	res = new(pms.BannerViewRes)
	res.PmsBannerViewModel = data
	return
}
func (c *ControllerPms) BannerEdit(ctx context.Context, req *pms.BannerEditReq) (res *pms.BannerEditRes, err error) {
	err = service.BasicsBanner().Edit(ctx, &req.PmsBannerEditInp)
	return
}
func (c *ControllerPms) BannerDelete(ctx context.Context, req *pms.BannerDeleteReq) (res *pms.BannerDeleteRes, err error) {
	err = service.BasicsBanner().Delete(ctx, &req.PmsBannerDeleteInp)
	return
}
