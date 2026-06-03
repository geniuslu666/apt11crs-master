package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) ActivityList(ctx context.Context, req *food.ActivityListReq) (res *food.ActivityListRes, err error) {
	list, totalCount, err := service.FoodActivity().List(ctx, &req.FoodActivityListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodActivityListModel{}
	}

	res = new(food.ActivityListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) ActivityExport(ctx context.Context, req *food.ActivityExportReq) (res *food.ActivityExportRes, err error) {
	err = service.FoodActivity().Export(ctx, &req.FoodActivityListInp)
	return
}
func (c *ControllerFood) ActivityView(ctx context.Context, req *food.ActivityViewReq) (res *food.ActivityViewRes, err error) {
	data, err := service.FoodActivity().View(ctx, &req.FoodActivityViewInp)
	if err != nil {
		return
	}

	res = new(food.ActivityViewRes)
	res.FoodActivityViewModel = data
	return
}
func (c *ControllerFood) ActivityEdit(ctx context.Context, req *food.ActivityEditReq) (res *food.ActivityEditRes, err error) {
	err = service.FoodActivity().Edit(ctx, &req.FoodActivityEditInp)
	return
}
func (c *ControllerFood) ActivityDelete(ctx context.Context, req *food.ActivityDeleteReq) (res *food.ActivityDeleteRes, err error) {
	err = service.FoodActivity().Delete(ctx, &req.FoodActivityDeleteInp)
	return
}
func (c *ControllerFood) ActivityStatus(ctx context.Context, req *food.ActivityStatusReq) (res *food.ActivityStatusRes, err error) {
	err = service.FoodActivity().Status(ctx, &req.FoodActivityStatusInp)
	return
}
