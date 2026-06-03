package admin

import (
	"APT/api/admin/travel"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
)

func (c *ControllerTravel) ProductList(ctx context.Context, req *travel.ProductListReq) (res *travel.ProductListRes, err error) {
	list, totalCount, err := service.TravelProduct().List(ctx, &req.TravelProductListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelProductListModel{}
	}
	res = new(travel.ProductListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerTravel) ProductView(ctx context.Context, req *travel.ProductViewReq) (res *travel.ProductViewRes, err error) {
	data, err := service.TravelProduct().View(ctx, &req.TravelProductViewInp)
	if err != nil {
		return
	}
	res = new(travel.ProductViewRes)
	res.TravelProductViewModel = data
	return
}

func (c *ControllerTravel) ProductEdit(ctx context.Context, req *travel.ProductEditReq) (res *travel.ProductEditRes, err error) {
	err = service.TravelProduct().Edit(ctx, &req.TravelProductEditInp)
	return
}

func (c *ControllerTravel) ProductDelete(ctx context.Context, req *travel.ProductDeleteReq) (res *travel.ProductDeleteRes, err error) {
	err = service.TravelProduct().Delete(ctx, &req.TravelProductDeleteInp)
	return
}

func (c *ControllerTravel) ProductStatus(ctx context.Context, req *travel.ProductStatusReq) (res *travel.ProductStatusRes, err error) {
	err = service.TravelProduct().Status(ctx, &req.TravelProductStatusInp)
	return
}

func (c *ControllerTravel) ProductRecycleList(ctx context.Context, req *travel.ProductRecycleListReq) (res *travel.ProductRecycleListRes, err error) {
	list, totalCount, err := service.TravelProduct().RecycleList(ctx, &req.TravelProductListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelProductListModel{}
	}
	res = new(travel.ProductRecycleListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerTravel) ProductRestore(ctx context.Context, req *travel.ProductRestoreReq) (res *travel.ProductRestoreRes, err error) {
	err = service.TravelProduct().Restore(ctx, &req.TravelProductDeleteInp)
	return
}
