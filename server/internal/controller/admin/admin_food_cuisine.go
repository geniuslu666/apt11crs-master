package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) CuisineList(ctx context.Context, req *food.CuisineListReq) (res *food.CuisineListRes, err error) {
	list, totalCount, err := service.FoodCuisine().List(ctx, &req.FoodCuisineListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodCuisineListModel{}
	}

	res = new(food.CuisineListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) CuisineView(ctx context.Context, req *food.CuisineViewReq) (res *food.CuisineViewRes, err error) {
	data, err := service.FoodCuisine().View(ctx, &req.FoodCuisineViewInp)
	if err != nil {
		return
	}

	res = new(food.CuisineViewRes)
	res.FoodCuisineViewModel = data
	return
}
func (c *ControllerFood) CuisineEdit(ctx context.Context, req *food.CuisineEditReq) (res *food.CuisineEditRes, err error) {
	err = service.FoodCuisine().Edit(ctx, &req.FoodCuisineEditInp)
	return
}
func (c *ControllerFood) CuisineDelete(ctx context.Context, req *food.CuisineDeleteReq) (res *food.CuisineDeleteRes, err error) {
	err = service.FoodCuisine().Delete(ctx, &req.FoodCuisineDeleteInp)
	return
}
func (c *ControllerFood) CuisineMaxSort(ctx context.Context, req *food.CuisineMaxSortReq) (res *food.CuisineMaxSortRes, err error) {
	data, err := service.FoodCuisine().MaxSort(ctx, &req.FoodCuisineMaxSortInp)
	if err != nil {
		return
	}

	res = new(food.CuisineMaxSortRes)
	res.FoodCuisineMaxSortModel = data
	return
}
func (c *ControllerFood) CuisineStatus(ctx context.Context, req *food.CuisineStatusReq) (res *food.CuisineStatusRes, err error) {
	err = service.FoodCuisine().Status(ctx, &req.FoodCuisineStatusInp)
	return
}
