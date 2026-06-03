package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) CooperateTypeList(ctx context.Context, req *food.CooperateTypeListReq) (res *food.CooperateTypeListRes, err error) {
	list, totalCount, err := service.FoodCooperateType().List(ctx, &req.FoodCooperateTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodCooperateTypeListModel{}
	}

	res = new(food.CooperateTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) CooperateTypeExport(ctx context.Context, req *food.CooperateTypeExportReq) (res *food.CooperateTypeExportRes, err error) {
	err = service.FoodCooperateType().Export(ctx, &req.FoodCooperateTypeListInp)
	return
}
func (c *ControllerFood) CooperateTypeView(ctx context.Context, req *food.CooperateTypeViewReq) (res *food.CooperateTypeViewRes, err error) {
	data, err := service.FoodCooperateType().View(ctx, &req.FoodCooperateTypeViewInp)
	if err != nil {
		return
	}

	res = new(food.CooperateTypeViewRes)
	res.FoodCooperateTypeViewModel = data
	return
}
func (c *ControllerFood) CooperateTypeEdit(ctx context.Context, req *food.CooperateTypeEditReq) (res *food.CooperateTypeEditRes, err error) {
	err = service.FoodCooperateType().Edit(ctx, &req.FoodCooperateTypeEditInp)
	return
}
func (c *ControllerFood) CooperateTypeDelete(ctx context.Context, req *food.CooperateTypeDeleteReq) (res *food.CooperateTypeDeleteRes, err error) {
	err = service.FoodCooperateType().Delete(ctx, &req.FoodCooperateTypeDeleteInp)
	return
}
func (c *ControllerFood) CooperateTypeStatus(ctx context.Context, req *food.CooperateTypeStatusReq) (res *food.CooperateTypeStatusRes, err error) {
	err = service.FoodCooperateType().Status(ctx, &req.FoodCooperateTypeStatusInp)
	return
}
