package admin

import (
	"APT/api/admin/car"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
)

func (c *ControllerCar) BannerList(ctx context.Context, req *car.BannerListReq) (res *car.BannerListRes, err error) {
	list, totalCount, err := service.CarBanner().List(ctx, &req.CarBannerListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarBannerListModel{}
	}

	res = new(car.BannerListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) BannerView(ctx context.Context, req *car.BannerViewReq) (res *car.BannerViewRes, err error) {
	data, err := service.CarBanner().View(ctx, &req.CarBannerViewInp)
	if err != nil {
		return
	}

	res = new(car.BannerViewRes)
	res.CarBannerViewModel = data
	return
}
func (c *ControllerCar) BannerEdit(ctx context.Context, req *car.BannerEditReq) (res *car.BannerEditRes, err error) {
	err = service.CarBanner().Edit(ctx, &req.CarBannerEditInp)
	return
}
func (c *ControllerCar) BannerDelete(ctx context.Context, req *car.BannerDeleteReq) (res *car.BannerDeleteRes, err error) {
	err = service.CarBanner().Delete(ctx, &req.CarBannerDeleteInp)
	return
}
func (c *ControllerCar) BannerStatus(ctx context.Context, req *car.BannerStatusReq) (res *car.BannerStatusRes, err error) {
	err = service.CarBanner().Status(ctx, &req.CarBannerStatusInp)
	return
}
func (c *ControllerCar) BannerMaxSort(ctx context.Context, req *car.BannerMaxSortReq) (res *car.BannerMaxSortRes, err error) {
	data, err := service.CarBanner().MaxSort(ctx, &req.CarBannerMaxSortInp)
	if err != nil {
		return
	}

	res = new(car.BannerMaxSortRes)
	res.CarBannerMaxSortModel = data
	return
}
