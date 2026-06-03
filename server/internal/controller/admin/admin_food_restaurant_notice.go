package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) RestaurantNoticeList(ctx context.Context, req *food.RestaurantNoticeListReq) (res *food.RestaurantNoticeListRes, err error) {
	list, totalCount, err := service.FoodRestaurantNotice().List(ctx, &req.FoodRestaurantNoticeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodRestaurantNoticeListModel{}
	}

	res = new(food.RestaurantNoticeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) RestaurantNoticeExport(ctx context.Context, req *food.RestaurantNoticeExportReq) (res *food.RestaurantNoticeExportRes, err error) {
	err = service.FoodRestaurantNotice().Export(ctx, &req.FoodRestaurantNoticeListInp)
	return
}
func (c *ControllerFood) RestaurantNoticeView(ctx context.Context, req *food.RestaurantNoticeViewReq) (res *food.RestaurantNoticeViewRes, err error) {
	data, err := service.FoodRestaurantNotice().View(ctx, &req.FoodRestaurantNoticeViewInp)
	if err != nil {
		return
	}

	res = new(food.RestaurantNoticeViewRes)
	res.FoodRestaurantNoticeViewModel = data
	return
}
func (c *ControllerFood) RestaurantNoticeEdit(ctx context.Context, req *food.RestaurantNoticeEditReq) (res *food.RestaurantNoticeEditRes, err error) {
	err = service.FoodRestaurantNotice().Edit(ctx, &req.FoodRestaurantNoticeEditInp)
	return
}
func (c *ControllerFood) RestaurantNoticeDelete(ctx context.Context, req *food.RestaurantNoticeDeleteReq) (res *food.RestaurantNoticeDeleteRes, err error) {
	err = service.FoodRestaurantNotice().Delete(ctx, &req.FoodRestaurantNoticeDeleteInp)
	return
}
func (c *ControllerFood) RestaurantNoticeMaxSort(ctx context.Context, req *food.RestaurantNoticeMaxSortReq) (res *food.RestaurantNoticeMaxSortRes, err error) {
	data, err := service.FoodRestaurantNotice().MaxSort(ctx, &req.FoodRestaurantNoticeMaxSortInp)
	if err != nil {
		return
	}

	res = new(food.RestaurantNoticeMaxSortRes)
	res.FoodRestaurantNoticeMaxSortModel = data
	return
}
func (c *ControllerFood) RestaurantNoticeStatus(ctx context.Context, req *food.RestaurantNoticeStatusReq) (res *food.RestaurantNoticeStatusRes, err error) {
	err = service.FoodRestaurantNotice().Status(ctx, &req.FoodRestaurantNoticeStatusInp)
	return
}
