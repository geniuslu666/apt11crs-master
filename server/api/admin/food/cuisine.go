package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type CuisineListReq struct {
	g.Meta `path:"/foodCuisine/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅菜系列表"`
	input_food.FoodCuisineListInp
}

type CuisineListRes struct {
	input_form.PageRes
	List []*input_food.FoodCuisineListModel `json:"list"   dc:"数据列表"`
}

type CuisineViewReq struct {
	g.Meta `path:"/foodCuisine/view" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅菜系指定信息"`
	input_food.FoodCuisineViewInp
}

type CuisineViewRes struct {
	*input_food.FoodCuisineViewModel
}

type CuisineEditReq struct {
	g.Meta `path:"/foodCuisine/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增餐厅菜系"`
	input_food.FoodCuisineEditInp
}

type CuisineEditRes struct{}

type CuisineDeleteReq struct {
	g.Meta `path:"/foodCuisine/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除餐厅菜系"`
	input_food.FoodCuisineDeleteInp
}

type CuisineDeleteRes struct{}

type CuisineMaxSortReq struct {
	g.Meta `path:"/foodCuisine/maxSort" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅菜系最大排序"`
	input_food.FoodCuisineMaxSortInp
}

type CuisineMaxSortRes struct {
	*input_food.FoodCuisineMaxSortModel
}

type CuisineStatusReq struct {
	g.Meta `path:"/foodCuisine/status" method:"post" tags:"ADMIN_FOOD" summary:"更新餐厅菜系状态"`
	input_food.FoodCuisineStatusInp
}

type CuisineStatusRes struct{}
