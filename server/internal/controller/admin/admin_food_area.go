package admin

import (
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) AreaList(ctx context.Context, req *food.AreaListReq) (res *food.AreaListRes, err error) {
	list, totalCount, err := service.FoodArea().List(ctx, &req.FoodAreaListInp)
	if err != nil {
		return
	}

	res = new(food.AreaListRes)
	res.FoodAreaListModel = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) AreaView(ctx context.Context, req *food.AreaViewReq) (res *food.AreaViewRes, err error) {
	data, err := service.FoodArea().View(ctx, &req.FoodAreaViewInp)
	if err != nil {
		return
	}

	res = new(food.AreaViewRes)
	res.FoodAreaViewModel = data
	return
}
func (c *ControllerFood) AreaEdit(ctx context.Context, req *food.AreaEditReq) (res *food.AreaEditRes, err error) {
	err = service.FoodArea().Edit(ctx, &req.FoodAreaEditInp)
	return
}
func (c *ControllerFood) AreaDelete(ctx context.Context, req *food.AreaDeleteReq) (res *food.AreaDeleteRes, err error) {
	err = service.FoodArea().Delete(ctx, &req.FoodAreaDeleteInp)
	return
}
func (c *ControllerFood) AreaStatus(ctx context.Context, req *food.AreaStatusReq) (res *food.AreaStatusRes, err error) {
	err = service.FoodArea().Status(ctx, &req.FoodAreaStatusInp)
	return
}
