package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) GoodsList(ctx context.Context, req *food.GoodsListReq) (res *food.GoodsListRes, err error) {
	list, totalCount, err := service.FoodGoods().List(ctx, &req.FoodGoodsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodGoodsListModel{}
	}

	res = new(food.GoodsListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) GoodsExport(ctx context.Context, req *food.GoodsExportReq) (res *food.GoodsExportRes, err error) {
	err = service.FoodGoods().Export(ctx, &req.FoodGoodsListInp)
	return
}
func (c *ControllerFood) GoodsView(ctx context.Context, req *food.GoodsViewReq) (res *food.GoodsViewRes, err error) {
	data, err := service.FoodGoods().View(ctx, &req.FoodGoodsViewInp)
	if err != nil {
		return
	}

	res = new(food.GoodsViewRes)
	res.FoodGoodsViewModel = data
	return
}
func (c *ControllerFood) GoodsEdit(ctx context.Context, req *food.GoodsEditReq) (res *food.GoodsEditRes, err error) {
	err = service.FoodGoods().Edit(ctx, &req.FoodGoodsEditInp)
	return
}
func (c *ControllerFood) GoodsDelete(ctx context.Context, req *food.GoodsDeleteReq) (res *food.GoodsDeleteRes, err error) {
	err = service.FoodGoods().Delete(ctx, &req.FoodGoodsDeleteInp)
	return
}
func (c *ControllerFood) GoodsStatus(ctx context.Context, req *food.GoodsStatusReq) (res *food.GoodsStatusRes, err error) {
	err = service.FoodGoods().Status(ctx, &req.FoodGoodsStatusInp)
	return
}
