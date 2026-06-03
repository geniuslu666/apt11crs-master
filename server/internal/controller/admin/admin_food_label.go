package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) LabelList(ctx context.Context, req *food.LabelListReq) (res *food.LabelListRes, err error) {
	list, totalCount, err := service.FoodLabel().List(ctx, &req.FoodLabelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodLabelListModel{}
	}

	res = new(food.LabelListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) LabelView(ctx context.Context, req *food.LabelViewReq) (res *food.LabelViewRes, err error) {
	data, err := service.FoodLabel().View(ctx, &req.FoodLabelViewInp)
	if err != nil {
		return
	}

	res = new(food.LabelViewRes)
	res.FoodLabelViewModel = data
	return
}
func (c *ControllerFood) LabelEdit(ctx context.Context, req *food.LabelEditReq) (res *food.LabelEditRes, err error) {
	err = service.FoodLabel().Edit(ctx, &req.FoodLabelEditInp)
	return
}
func (c *ControllerFood) LabelDelete(ctx context.Context, req *food.LabelDeleteReq) (res *food.LabelDeleteRes, err error) {
	err = service.FoodLabel().Delete(ctx, &req.FoodLabelDeleteInp)
	return
}
func (c *ControllerFood) LabelMaxSort(ctx context.Context, req *food.LabelMaxSortReq) (res *food.LabelMaxSortRes, err error) {
	data, err := service.FoodLabel().MaxSort(ctx, &req.FoodLabelMaxSortInp)
	if err != nil {
		return
	}

	res = new(food.LabelMaxSortRes)
	res.FoodLabelMaxSortModel = data
	return
}
func (c *ControllerFood) LabelStatus(ctx context.Context, req *food.LabelStatusReq) (res *food.LabelStatusRes, err error) {
	err = service.FoodLabel().Status(ctx, &req.FoodLabelStatusInp)
	return
}
