package app

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"

	"APT/api/app/travel"
)

func (c *ControllerTravel) ProductList(ctx context.Context, req *travel.ProductListReq) (res *travel.ProductListRes, err error) {
	res = new(travel.ProductListRes)
	if res.List, res.Count, err = service.TravelProduct().AppList(ctx, &input_travel.TravelProductAppListInp{
		PageReq: input_form.PageReq{
			Page:       req.PageNum,
			PerPage:    req.PageSize,
			Pagination: true,
		},
	}); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) ProductView(ctx context.Context, req *travel.ProductViewReq) (res *travel.ProductViewRes, err error) {
	res = new(travel.ProductViewRes)
	if res.TravelProductAppViewModel, err = service.TravelProduct().AppView(ctx, &req.TravelProductAppViewInp); err != nil {
		return
	}
	return
}

func (c *ControllerTravel) ProductSkuStock(ctx context.Context, req *travel.ProductSkuStockReq) (res *travel.ProductSkuStockRes, err error) {
	res = new(travel.ProductSkuStockRes)
	if res.List, err = service.TravelProduct().SkuStock(ctx, &req.TravelProductSkuStockInp); err != nil {
		return
	}
	return
}
