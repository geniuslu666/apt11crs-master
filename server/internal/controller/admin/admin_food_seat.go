package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) SeatList(ctx context.Context, req *food.SeatListReq) (res *food.SeatListRes, err error) {
	list, totalCount, err := service.FoodSeat().List(ctx, &req.FoodSeatListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodSeatListModel{}
	}

	res = new(food.SeatListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) SeatView(ctx context.Context, req *food.SeatViewReq) (res *food.SeatViewRes, err error) {
	data, err := service.FoodSeat().View(ctx, &req.FoodSeatViewInp)
	if err != nil {
		return
	}

	res = new(food.SeatViewRes)
	res.FoodSeatViewModel = data
	return
}
func (c *ControllerFood) SeatEdit(ctx context.Context, req *food.SeatEditReq) (res *food.SeatEditRes, err error) {
	err = service.FoodSeat().Edit(ctx, &req.FoodSeatEditInp)
	return
}
func (c *ControllerFood) SeatDelete(ctx context.Context, req *food.SeatDeleteReq) (res *food.SeatDeleteRes, err error) {
	err = service.FoodSeat().Delete(ctx, &req.FoodSeatDeleteInp)
	return
}
func (c *ControllerFood) SeatStatus(ctx context.Context, req *food.SeatStatusReq) (res *food.SeatStatusRes, err error) {
	err = service.FoodSeat().Status(ctx, &req.FoodSeatStatusInp)
	return
}
