package admin

import (
	"APT/api/admin/food"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"
)

func (c *ControllerFood) RestaurantList(ctx context.Context, req *food.RestaurantListReq) (res *food.RestaurantListRes, err error) {
	list, totalCount, err := service.FoodRestaurant().List(ctx, &req.FoodRestaurantListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodRestaurantListModel{}
	}

	res = new(food.RestaurantListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) RestaurantAllList(ctx context.Context, req *food.RestaurantAllListReq) (res *food.RestaurantAllListRes, err error) {
	list, err := service.FoodRestaurant().All(ctx, &req.FoodRestaurantListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodRestaurantAllListModel{}
	}

	res = new(food.RestaurantAllListRes)
	res.List = list
	return
}
func (c *ControllerFood) RestaurantView(ctx context.Context, req *food.RestaurantViewReq) (res *food.RestaurantViewRes, err error) {
	data, err := service.FoodRestaurant().View(ctx, &req.FoodRestaurantViewInp)
	if err != nil {
		return
	}

	res = new(food.RestaurantViewRes)
	res.FoodRestaurantViewModel = data
	return
}
func (c *ControllerFood) RestaurantEdit(ctx context.Context, req *food.RestaurantEditReq) (res *food.RestaurantEditRes, err error) {
	err = service.FoodRestaurant().Edit(ctx, &req.FoodRestaurantEditInp)
	return
}
func (c *ControllerFood) RestaurantDelete(ctx context.Context, req *food.RestaurantDeleteReq) (res *food.RestaurantDeleteRes, err error) {
	err = service.FoodRestaurant().Delete(ctx, &req.FoodRestaurantDeleteInp)
	return
}
func (c *ControllerFood) RestaurantStatus(ctx context.Context, req *food.RestaurantStatusReq) (res *food.RestaurantStatusRes, err error) {
	err = service.FoodRestaurant().Status(ctx, &req.FoodRestaurantStatusInp)
	return
}
func (c *ControllerFood) RestaurantResetCode(ctx context.Context, req *food.RestaurantResetCodeReq) (res *food.RestaurantResetCodeRes, err error) {
	err = service.FoodRestaurant().ResetVerifyCode(ctx, &req.FoodRestaurantRestVerifyCodeInp)
	return

}
func (c *ControllerFood) RestaurantSwitch(ctx context.Context, req *food.RestaurantSwitchReq) (res *food.RestaurantSwitchRes, err error) {
	err = service.FoodRestaurant().Switch(ctx, &req.FoodRestaurantSwitchInp)
	return
}
func (c *ControllerFood) RestaurantSort(ctx context.Context, req *food.RestaurantSortReq) (res *food.RestaurantSortRes, err error) {
	err = service.FoodRestaurant().RestaurantSort(ctx, &req.FoodRestaurantSortInp)
	return
}
