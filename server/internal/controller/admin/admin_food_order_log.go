package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) OrderLogList(ctx context.Context, req *food.OrderLogListReq) (res *food.OrderLogListRes, err error) {
	list, totalCount, err := service.FoodOrderLog().List(ctx, &req.FoodOrderLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodOrderLogListModel{}
	}

	res = new(food.OrderLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
