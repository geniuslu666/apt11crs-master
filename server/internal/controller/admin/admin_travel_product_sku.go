package admin

import (
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"

	"APT/api/admin/travel"
)

func (c *ControllerTravel) ProductSkuList(ctx context.Context, req *travel.ProductSkuListReq) (res *travel.ProductSkuListRes, err error) {
	list, totalCount, err := service.TravelProductSku().List(ctx, &req.TravelProductSkuListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelProductSkuListModel{}
	}
	res = new(travel.ProductSkuListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTravel) ProductSkuView(ctx context.Context, req *travel.ProductSkuViewReq) (res *travel.ProductSkuViewRes, err error) {
	data, err := service.TravelProductSku().View(ctx, &req.TravelProductSkuViewInp)
	if err != nil {
		return
	}
	res = new(travel.ProductSkuViewRes)
	res.TravelProductSkuViewModel = data
	return
}
func (c *ControllerTravel) ProductSkuEdit(ctx context.Context, req *travel.ProductSkuEditReq) (res *travel.ProductSkuEditRes, err error) {
	err = service.TravelProductSku().Edit(ctx, &req.TravelProductSkuEditInp)
	return
}
func (c *ControllerTravel) ProductSkuDelete(ctx context.Context, req *travel.ProductSkuDeleteReq) (res *travel.ProductSkuDeleteRes, err error) {
	err = service.TravelProductSku().Delete(ctx, &req.TravelProductSkuDeleteInp)
	return
}
func (c *ControllerTravel) ProductSkuStatus(ctx context.Context, req *travel.ProductSkuStatusReq) (res *travel.ProductSkuStatusRes, err error) {
	err = service.TravelProductSku().Status(ctx, &req.TravelProductSkuStatusInp)
	return
}
