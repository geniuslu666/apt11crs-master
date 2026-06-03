package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) MaintenanceList(ctx context.Context, req *food.MaintenanceListReq) (res *food.MaintenanceListRes, err error) {
	list, totalCount, err := service.FoodMaintenance().List(ctx, &req.FoodMaintenanceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodMaintenanceListModel{}
	}

	res = new(food.MaintenanceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) MaintenanceView(ctx context.Context, req *food.MaintenanceViewReq) (res *food.MaintenanceViewRes, err error) {
	data, err := service.FoodMaintenance().View(ctx, &req.FoodMaintenanceViewInp)
	if err != nil {
		return
	}

	res = new(food.MaintenanceViewRes)
	res.FoodMaintenanceViewModel = data
	return
}
func (c *ControllerFood) MaintenanceEdit(ctx context.Context, req *food.MaintenanceEditReq) (res *food.MaintenanceEditRes, err error) {
	err = service.FoodMaintenance().Edit(ctx, &req.FoodMaintenanceEditInp)
	return
}
func (c *ControllerFood) MaintenanceDelete(ctx context.Context, req *food.MaintenanceDeleteReq) (res *food.MaintenanceDeleteRes, err error) {
	err = service.FoodMaintenance().Delete(ctx, &req.FoodMaintenanceDeleteInp)
	return
}
func (c *ControllerFood) MaintenanceLanguageList(ctx context.Context, req *food.MaintenanceLanguageListReq) (res *food.MaintenanceLanguageListRes, err error) {
	list, err := service.FoodMaintenance().LanguageList(ctx, &req.FoodMaintenanceLanguageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodMaintenanceLanguageListModel{}
	}

	res = new(food.MaintenanceLanguageListRes)
	res.List = list
	return
}
