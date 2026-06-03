package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type RestaurantListReq struct {
	g.Meta `path:"/foodRestaurant/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅管理列表"`
	input_food.FoodRestaurantListInp
}

type RestaurantListRes struct {
	input_form.PageRes
	List []*input_food.FoodRestaurantListModel `json:"list"   dc:"数据列表"`
}

type RestaurantAllListReq struct {
	g.Meta `path:"/foodRestaurant/all" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅列表"`
	input_food.FoodRestaurantListInp
}

type RestaurantAllListRes struct {
	List []*input_food.FoodRestaurantAllListModel `json:"list"   dc:"所有餐厅-数据列表"`
}

type RestaurantViewReq struct {
	g.Meta `path:"/foodRestaurant/view" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅管理指定信息"`
	input_food.FoodRestaurantViewInp
}

type RestaurantViewRes struct {
	*input_food.FoodRestaurantViewModel
}

type RestaurantEditReq struct {
	g.Meta `path:"/foodRestaurant/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增餐厅管理"`
	input_food.FoodRestaurantEditInp
}

type RestaurantEditRes struct{}

type RestaurantDeleteReq struct {
	g.Meta `path:"/foodRestaurant/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除餐厅管理"`
	input_food.FoodRestaurantDeleteInp
}

type RestaurantDeleteRes struct{}

type RestaurantStatusReq struct {
	g.Meta `path:"/foodRestaurant/status" method:"post" tags:"ADMIN_FOOD" summary:"更新餐厅状态"`
	input_food.FoodRestaurantStatusInp
}

type RestaurantStatusRes struct{}

type RestaurantResetCodeReq struct {
	g.Meta `path:"/foodRestaurant/resetVerifyCode" method:"post" tags:"ADMIN_FOOD" summary:"餐厅_重置核销码"`
	input_food.FoodRestaurantRestVerifyCodeInp
}

type RestaurantResetCodeRes struct{}

type RestaurantSwitchReq struct {
	g.Meta `path:"/foodRestaurant/switch" method:"post" tags:"ADMIN_FOOD" summary:"餐厅_更新预定状态"`
	input_food.FoodRestaurantSwitchInp
}

type RestaurantSwitchRes struct{}

// RestaurantSortReq 餐厅排序
type RestaurantSortReq struct {
	g.Meta `path:"/foodRestaurant/sortUpdate" method:"post" tags:"ADMIN_FOOD" summary:"餐厅_排序"`
	input_food.FoodRestaurantSortInp
}

type RestaurantSortRes struct{}
